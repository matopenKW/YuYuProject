package main

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"encoding/json"
	"log"
)

func main() {

	teamList, err := getTeamJson()
	if err != nil {
		log.Fatalln(err)
	}

	teamRateList, err := getTeamRateJson()
	if err != nil {
		log.Fatalln(err)
	}

	var barList []*dto.TeamRate
	for _, team := range teamList {
		for _, teamRate := range teamRateList {
			if team.Id == teamRate.Id {
				all := teamRate
				all.All = (teamRate.West + teamRate.East) / 2
				barList = append(barList, all)
			}
		}
	}

	for _, v := range barList {
		log.Println(v)
	}

}

func getTeamJson() ([]*dto.Team, error) {
	bytes, err := util.ReadFile("json/team.json")
	if err != nil {
		return nil, err
	}

	var teamList []*dto.Team
	err = json.Unmarshal(bytes, &teamList)
	if err != nil {
		return nil, err
	}

	return teamList, nil
}

func getTeamRateJson() ([]*dto.TeamRate, error) {
	bytes, err := util.ReadFile("json/teamRate.json")
	if err != nil {
		return nil, err
	}

	var teamRateList []*dto.TeamRate
	err = json.Unmarshal(bytes, &teamRateList)
	if err != nil {
		return nil, err
	}

	return teamRateList, nil
}
