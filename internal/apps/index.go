package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/enum"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowMainPage(ctx *gin.Context) (map[string]interface{}, error) {

	session := sessions.Default(ctx)
	userId := session.Get("userId")
	team := getUser(userId)

	teamIdView := true
	if team == nil {
		teamIdView = false
	}

	teamDao := dao.GetTeamDao()
	teamList, err := teamDao()
	if err != nil {
		return nil, err
	}

	list, err := getBarList(teamList)
	if err != nil {
		return nil, err
	}

	productMap, err := getProductBarList(teamList)
	if err != nil {
		return nil, err
	}

	log.Println("productList :", productMap)

	return map[string]interface{}{
		"list":         list,
		"productList1": productMap[enum.Nitori],
		"productList2": productMap[enum.Daiso],
		"productList3": productMap[enum.Hands],
		"teamIdView":   teamIdView,
		"team":         team,
	}, nil
}

func getBarList(teamList []*dto.Team) ([]*dto.Team, error) {

	var westCnt, eastCnt int
	var barList []*dto.Team
	for _, team := range teamList {
		bar := team
		bar.ClassName = team.ClassName
		bar.All = (team.West + team.East) / 2
		barList = append(barList, bar)

		westCnt += team.West
		eastCnt += team.East
	}

	bar := &dto.Team{}
	bar.Id = "N"
	bar.ClassName = "none_team"
	bar.West = 100 - westCnt
	bar.East = 100 - eastCnt
	bar.All = (bar.West + bar.East) / 2
	barList = append(barList, bar)

	return barList, nil
}

func getProductBarList(teamList []*dto.Team) (map[string][]*dto.TeamProduct, error) {

	productBarMap := make(map[string][]*dto.TeamProduct)

	tenantIdList := [3]string{enum.Nitori, enum.Daiso, enum.Hands}

	getProductDao := dao.GetProductDao()
	for _, tenantId := range tenantIdList {

		var productBarList []*dto.TeamProduct
		var total int
		if tenantId == enum.Daiso {
			total = 15
		} else {
			total = 30
		}

		for _, v := range teamList {
			productList, err := getProductDao(v.Id, tenantId)
			if err != nil {
				return nil, err
			}
			productNum := len(productList)
			rate := (productNum * 100) / total
			teamProduct := &dto.TeamProduct{
				v,
				tenantId,
				productList,
				productNum,
				rate,
			}

			productBarList = append(productBarList, teamProduct)

		}

		productBarMap[tenantId] = productBarList
	}
	return productBarMap, nil
}

// 画面描画はセッション切れててもOK
func getUser(userId interface{}) *dto.Team {
	log.Println("getUser start")
	defer log.Println("getUser end")

	if userId == nil {
		log.Println("UserId Is null")
		return nil
	}
	getSingleTeamDao := dao.GetSingleTeamDao()
	team, err := getSingleTeamDao(userId.(string))

	if err != nil {
		log.Println(err)
		return nil
	}
	return team
}
