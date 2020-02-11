package main

import (
	"YuYuProject/internal/apps"
	"YuYuProject/pkg/db"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	// setting session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// ログインページ表示
	router.GET("/", viewLogin)
	router.POST("/", viewLogin)
	router.GET("/login", viewLogin)
	router.POST("/login", viewLogin)

	// ログイン処理
	router.POST("/login:cmd/login", login)

	// HOME画面
	router.GET("/index", index)

	// データ取得処理
	router.POST("/floor", showFloor)
	router.GET("/registSerial", checkSession(registSerial))
	router.GET("/ragistProduct", checkSession(ragistProduct))

	// WebAPI
	router.GET("/updateTenantoTeam", updateTenantoTeam)

	router.Run()
}

func viewLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func login(ctx *gin.Context) {
	err := apps.Login(ctx)
	if err != nil {
		log.Println(err)
		ctx.HTML(http.StatusInternalServerError, "505.html", gin.H{})
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	}
}

func index(ctx *gin.Context) {

	form, err := apps.ShowMainPage(ctx)

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "505.html", gin.H{})
	} else {
		ctx.HTML(http.StatusOK, "index.html", form)
	}
}

func checkSession(proc func(ctx *gin.Context, userRec *auth.UserRecord)) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		timeOutFunc := func(ctx *gin.Context) {
			ctx.JSON(http.StatusRequestTimeout, createJsonMessage("セッションが切れました。ログインしなおしてください。"))
		}

		session := sessions.Default(ctx)
		userId := session.Get("userId")
		log.Printf("userId : %v\n", userId)

		authClient, err := db.OpenAuth()
		if err != nil {
			timeOutFunc(ctx)
		}

		// sessionのチェック
		var userRec *auth.UserRecord
		if userId != nil {
			userRec, err = db.GetUserRecord(authClient, userId.(string))
			if err != nil {
				log.Printf("error : %v\n", err)
				timeOutFunc(ctx)
			}
		} else {
			err := errors.New("session time out")
			log.Printf("error : %v\n", err)
			timeOutFunc(ctx)
		}

		// 処理実行
		proc(ctx, userRec)
	}
}

func showFloor(ctx *gin.Context) {

	res, err := apps.GetFloorData(ctx.Request)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, string(bytes))
	} else {
		ctx.JSON(200, string(bytes))
	}
}

func registSerial(ctx *gin.Context, userRec *auth.UserRecord) {

	err := apps.RegistSerial(userRec, ctx.Request)

	var satus int
	var msg string

	if err != nil {
		satus = http.StatusInternalServerError
		msg = err.Error()
	} else {
		satus = http.StatusOK
		msg = "成功しました。"
	}

	ctx.JSON(satus, createJsonMessage(msg))
}

func ragistProduct(ctx *gin.Context, userRec *auth.UserRecord) {
	err := apps.RegistProduct(userRec, ctx.Request)

	var satus int
	var msg string

	if err != nil {
		satus = http.StatusInternalServerError
		msg = err.Error()
	} else {
		satus = http.StatusOK
		msg = "成功しました。"
	}

	ctx.JSON(satus, createJsonMessage(msg))
}

func updateTenantoTeam(ctx *gin.Context) {
	err := apps.UpdateTenantoTeam()

	var satus int
	var msg string

	if err != nil {
		satus = http.StatusInternalServerError
		msg = err.Error()
	} else {
		satus = http.StatusOK
		msg = "成功しました。"
	}

	ctx.JSON(satus, createJsonMessage(msg))
}

func createJsonMessage(message string) string {
	bytes, err := json.Marshal(map[string]interface{}{
		"message": message,
	})
	if err != nil {
		return "{message: 'json marshal fail'}"
	} else {
		return string(bytes)
	}
}
