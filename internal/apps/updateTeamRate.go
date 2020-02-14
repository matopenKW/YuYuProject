package apps

import (
	"YuYuProject/internal/dao"
	"errors"
	"log"
)

func UpdateTeamRate() error {
	err := updateTeamRate()
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func updateTeamRate() error {

	log.Println("test")

	eastList := make([]string, 0, 0)
	eastList = append(eastList, "e_1", "e_2", "e_3", "e_4", "e_5", "e_6", "e_7")

	westList := make([]string, 0, 0)
	westList = append(westList, "w_1", "w_2", "w_3", "w_4", "w_5", "w_6", "w_7", "w_8")

	eastMap, eastScoreMap, err := getTeamRateMap(eastList, 162)
	if err != nil {
		return err
	}

	westMap, wastScoreMap, err := getTeamRateMap(westList, 58)
	if err != nil {
		return err
	}

	getTeanDao := dao.GetTeamDao()
	teamList, err := getTeanDao()
	if err != nil {
		return err
	}

	updateDao := dao.UpdateTeamDao()
	for _, v := range teamList {
		v.East = eastMap[v.Id]
		v.West = westMap[v.Id]
		v.Score = eastScoreMap[v.Id] + wastScoreMap[v.Id]

		err := updateDao(v)
		if err != nil {
			return err
		}

	}

	if false {
		return errors.New("")
	}

	return nil
}

func getTeamRateMap(floorIdList []string, totalScore int) (map[string]int, map[string]int, error) {
	teamMap := make(map[string]int)

	tenantDao := dao.GetTenatoDao()
	for _, floorId := range floorIdList {
		tenantList, err := tenantDao(floorId)
		if err != nil {
			return nil, nil, err
		}
		for _, tenant := range tenantList {
			score := tenant.Score
			if score < 1 {
				score = 1
			}

			teamMap[tenant.Acquisition] = teamMap[tenant.Acquisition] + score
		}
	}

	log.Println(teamMap)

	teamRateMap := make(map[string]int)
	for k, v := range teamMap {
		teamRateMap[k] = (v * 100 / totalScore)
	}
	return teamRateMap, teamMap, nil
}
