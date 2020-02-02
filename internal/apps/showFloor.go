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
	temaList, err := teamDao()
	if err != nil {
		return nil, err
	}

	teamMap := make(map[string]*dto.Team)
	cntMap := make(map[string]int)
	for _, v := range temaList {
		if v.Id == "" {
			v.Id = "N"
		}
		teamMap[v.Id] = v
		cntMap[v.Id] = 0
	}

	tenantoDao := dao.GetTenatoDao()
	floorList, err := tenantoDao(floorId)
	if err != nil {
		return nil, err
	}

	for _, floor := range floorList {
		if floor.Acquisition == "" {
			floor.Acquisition = "N"
		}
		floor.ClassName = teamMap[floor.Acquisition].ClassName
		cntMap[floor.Acquisition] = cntMap[floor.Acquisition] + 1
	}

	barList := make([]*dto.Team, 0, 0)
	for _, team := range temaList {
		var bar = &dto.Team{}
		bar.All = (cntMap[team.Id] * 100) / len(floorList)
		barList = append(barList, bar)
	}

	return map[string]interface{}{
		"barList":     barList,
		"tenantoList": floorList,
	}, nil
}
