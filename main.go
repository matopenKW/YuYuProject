package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", show("index"))
	router.GET("/modal", show("modal"))

	router.Run()
}

func show(htmlPath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.HTML(200, htmlPath+".html", gin.H{})
	}
}
