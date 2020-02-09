package main

import (
	"YuYuProject/internal/apps"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", index)
	router.POST("/floor", showFloor)
	router.GET("/registSerial", registSerial)
	router.GET("/ragistProduct", ragistProduct)

  router.GET("/updateTenantoTeam", updateTenantoTeam)

	router.Run()
}

func index(ctx *gin.Context) {

	form, err := apps.ShowMainPage()

	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "505.html", gin.H{})
	} else {
		ctx.HTML(http.StatusOK, "index.html", form)
	}
}

func showFloor(ctx *gin.Context) {

	res, err := apps.GetFloorData()
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

func registSerial(ctx *gin.Context) {

	err := apps.RegistSerial(ctx.Request)

	var satus int
	var msg string

	if err != nil {
		satus = http.StatusInternalServerError
		msg = err.Error()
	} else {
		satus = http.StatusOK
		msg = "成功しました。"
	}

	bytes, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{message: 'json marshal fail'}")
	} else {
		ctx.JSON(satus, string(bytes))
	}
}

func ragistProduct(ctx *gin.Context) {
	err := apps.RegistProduct(ctx.Request)
  
	var satus int
	var msg string

	if err != nil {
		satus = http.StatusInternalServerError
		msg = err.Error()
	} else {
		satus = http.StatusOK
		msg = "成功しました。"
	}

	bytes, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{message: 'json marshal fail'}")
	} else {
		ctx.JSON(satus, string(bytes))
	}
}

func updateTenantoTeam(ctx *gin.Context){
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

	bytes, err := json.Marshal(map[string]interface{}{
		"message": msg,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "{message: 'json marshal fail'}")
	} else {
		ctx.JSON(satus, string(bytes))
	}
}
