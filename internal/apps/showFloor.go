package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"log"
)

func GetFloorData() (map[string]interface{}, error) {

	res, err := getData("e_1")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func getData(floorId string) (map[string]interface{}, error) {

	teamDao := dao.GetTeamDao()
	teamList, err := teamDao()
	if err != nil {
		return nil, err
	}

	teamMap := make(map[string]*dto.Team)
	cntMap := make(map[string]int)
	for _, v := range teamList {
		teamMap[v.Id] = v
		cntMap[v.Id] = 0
	}

	tenantoDao := dao.GetTenatoDao()
	floorList, err := tenantoDao(floorId)
	if err != nil {
		return nil, err
	}

	for _, floor := range floorList {
		team := teamMap[floor.Acquisition]
		if team == nil {
			floor.Acquisition = "N"
			floor.ClassName = "none_team"
		} else {
			floor.ClassName = team.ClassName
		}
		cntMap[floor.Acquisition] = cntMap[floor.Acquisition] + 1
	}

	barList := make([]*dto.Team, 0, 0)
	sumRate := 0
	for _, team := range teamList {
		var bar = &dto.Team{}
		bar = team
		bar.All = (cntMap[team.Id] * 100) / len(floorList)
		sumRate += bar.All
		barList = append(barList, bar)
	}
	if sumRate != 100 {
		var bar = &dto.Team{}
		bar.Id = "N"
		bar.ClassName = "none_team"
		bar.All = 100 - sumRate
		barList = append(barList, bar)
	}
	
	return map[string]interface{}{
		"barList":     barList,
		"tenantoList": floorList,
	}, nil
}
