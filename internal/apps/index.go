package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
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

	list, err := getBarList()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"list":       list,
		"teamIdView": teamIdView,
		"team":       team,
	}, nil
}

func getBarList() ([]*dto.Team, error) {

	teamDao := dao.GetTeamDao()
	teamList, err := teamDao()
	if err != nil {
		return nil, err
	}

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

	for _, v := range barList {
		log.Println(v)
	}

	return barList, nil
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
