package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"log"
	"net/http"
)

func GetFloorData(req *http.Request) (map[string]interface{}, error) {

	req.ParseForm()
	if err := util.CheckNil(req.Form["floorId"], "フロアId"); err != nil {
		return nil, err
	}

	res, err := GetData(req.Form["floorId"][0])
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return res, nil
}

func GetData(floorId string) (map[string]interface{}, error) {

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
		if bar.All != 0 {
			barList = append(barList, bar)
		}
		sumRate += bar.All
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
