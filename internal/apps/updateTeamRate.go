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

	eastMap, err := getTeamRateMap(eastList)
	if err != nil {
		return err
	}
	log.Println(eastMap)
	westMap, err := getTeamRateMap(westList)
	if err != nil {
		return err
	}
	log.Println(westMap)

	if false {
		return errors.New("")
	}

	return nil
}

func getTeamRateMap(floorIdList []string) (map[string]int, error) {
	teamMap := make(map[string]int)

	tenantSize := 0
	tenantDao := dao.GetTenatoDao()
	for _, floorId := range floorIdList {
		tenantList, err := tenantDao(floorId)
		if err != nil {
			return nil, err
		}
		for _, tenant := range tenantList {
			teamMap[tenant.Acquisition] = teamMap[tenant.Acquisition] + 1
			tenantSize++
		}
	}

	teamRateMap := make(map[string]int)
	for k, v := range teamMap {
		teamRateMap[k] = (v * 100 / tenantSize)
	}
	return teamRateMap, nil

}

// func update() error {
// 	serialDao := dao.GetSerialDao()
// 	serials, err := serialDao()
// 	if err != nil {
// 		return err
// 	}

// 	tenantDao := dao.GetTenatoDao()
// 	for _, serial := range serials {
// 		errMsg := "Tenantoのレコードが不足しています FloorId:" + string(serial.FloorId) + " Seq:" + string(serial.Seq)

// 		if !serial.Acquired {
// 			continue
// 		}
// 		tenants, err := tenantDao(serial.FloorId)
// 		if err != nil {
// 			return err
// 		}
// 		if len(tenants) < serial.Seq {
// 			return errors.New(errMsg)
// 		}
// 		tenant := tenants[serial.Seq-1]
// 		if tenant.Seq != serial.Seq {
// 			return errors.New(errMsg)
// 		}

// 		tenant.Acquisition = serial.Acquisition
// 		updateTenanto := dao.UpdateTenantoDao()
// 		err = updateTenanto(serial.FloorId, tenant)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
