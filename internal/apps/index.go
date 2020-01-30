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

func getBarList() ([]*dto.TeamRate, error) {

	teamDao := dao.GetTeamDao()
	teamList, err := teamDao()
	if err != nil {
		return nil, err
	}

	teamRateDao := dao.GetTeamRateDao()
	teamRateList, err := teamRateDao()
	if err != nil {
		return nil, err
	}

	var westCnt, eastCnt int
	var barList []*dto.TeamRate
	for _, team := range teamList {
		for _, teamRate := range teamRateList {
			if team.Id == teamRate.Id {
				bar := teamRate
				bar.ClassName = team.ClassName
				bar.All = (teamRate.West + teamRate.East) / 2
				barList = append(barList, bar)

				westCnt += teamRate.West
				eastCnt += teamRate.East
			}
		}
	}

	bar := &dto.TeamRate{}
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
