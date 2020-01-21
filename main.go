package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Team struct {
	Name      string
	ClassName string
	Rate      int
}

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
		list, eastList, westList := getTeamList()

		fmt.Print(list)
		ctx.HTML(200, htmlPath+".html", gin.H{
			"allList":  list,
			"eastList": eastList,
			"westList": westList,
		})
	}
}

func fowardJson() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, "")
	}
}

func getTeamList() (allList, eastList, westList []*Team) {

	ALL := 30
	WEST := 15
	EAST := 15

	A_WEST := 3
	A_EAST := 4

	B_WEST := 3
	B_EAST := 6

	C_WEST := 9
	C_EAST := 5

	A := &Team{
		"A", "a_team", (A_WEST + A_EAST) * 100 / ALL,
	}

	B := &Team{
		"B", "b_team", (B_WEST + B_EAST) * 100 / ALL,
	}

	C := &Team{
		"C", "c_team", (C_WEST + C_EAST) * 100 / ALL,
	}
	allList = []*Team{A, B, C}

	for _, v := range allList {
		fmt.Println(v)
	}

	A = &Team{
		"A", "a_team", A_WEST * 100 / WEST,
	}

	B = &Team{
		"B", "b_team", B_WEST * 100 / WEST,
	}

	C = &Team{
		"C", "c_team", C_WEST * 100 / WEST,
	}
	eastList = []*Team{A, B, C}

	A = &Team{
		"A", "a_team", A_EAST * 100 / EAST,
	}

	B = &Team{
		"B", "b_team", B_EAST * 100 / EAST,
	}

	C = &Team{
		"C", "c_team", C_EAST * 100 / EAST,
	}
	westList = []*Team{A, B, C}

	return
}

func getRate(all int, num ...int) int {

	return 0
}
