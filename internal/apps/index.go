package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"log"
)

func ShowMainPage() (map[string]interface{}, error) {
	list, err := getBarList()
	if err != nil {
		return nil, err
	}

	log.Println(list)
	return map[string]interface{}{
		"list": list,
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
