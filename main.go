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
