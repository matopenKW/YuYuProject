package main

import (
	"YuYuProject/internal/dto"
	"encoding/json"
	"github.com/JustinTulloss/firebase"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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
	router.POST("/floor", fowardJson())

	router.Run()
}

func show(htmlPath string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		list, eastList, westList := getTeamList()

		log.Print(list)
		ctx.HTML(200, htmlPath+".html", gin.H{
			"allList":  list,
			"eastList": eastList,
			"westList": westList,
		})
	}
}

func fowardJson() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		res, err := getData()
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(505, "")
			return
		}
		log.Println(res)

		bytes, err := json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(505, string(bytes))
		} else {
			ctx.JSON(200, string(bytes))
		}

	}
}

func getList() []*dto.Tenanto {

	auth := "1zBUevXg2gtvtgCaqw3n591l0PR949ZzgwXabjKm"
	endpoint := "https://mac-001-1e4f9.firebaseio.com/"

	c := firebase.NewClient(endpoint+"/billdiing/e_1", auth, nil)

	retTenanto := func() interface{} {
		return &dto.Tenanto{}
	}

	var list []*dto.Tenanto
	for n := range c.Iterator(retTenanto) {
		dto := n.Value.(*dto.Tenanto)
		log.Println(dto)
		list = append(list, dto)
	}

	return list
}

func getListJson(tenanto string) ([]*dto.Tenanto, error) {
	bytes := useIoutilReadFile(tenanto)
	var tenantoList []*dto.Tenanto
	err := json.Unmarshal(bytes, &tenantoList)
	if err != nil {
		return nil, err
	}
	return tenantoList, nil
}

func getData() (map[string]interface{}, error) {
	list, err := getListJson("e_1")
	if err != nil {
		return nil, err
	}

	for _, v := range list {

		switch v.Acquisition {
		case "A":
			v.ClassName = "a_team"
		case "B":
			v.ClassName = "b_team"
		case "C":
			v.ClassName = "c_team"
		default:
			v.ClassName = "none-team"
		}
		log.Println(&v)
	}

	return map[string]interface{}{
		"JSON":        "Data",
		"tenantoList": list,
	}, nil
}

func useIoutilReadFile(tenanto string) []byte {
	bytes, err := ioutil.ReadFile("tool/json/" + tenanto + ".json")
	if err != nil {
		panic(err)
	}

	return bytes
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
		log.Println(v)
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
