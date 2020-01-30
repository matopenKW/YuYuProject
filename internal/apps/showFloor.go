package apps

import (
	"YuYuProject/internal/dto"
	"encoding/json"
	"github.com/JustinTulloss/firebase"
	"io/ioutil"
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
	list, err := getListJson(floorId)
	if err != nil {
		return nil, err
	}

	var aCnt, bCnt, cCnt int
	for _, v := range list {

		switch v.Acquisition {
		case "A":
			v.ClassName = "a_team"
			aCnt++
		case "B":
			v.ClassName = "b_team"
			bCnt++
		case "C":
			v.ClassName = "c_team"
			cCnt++
		default:
			v.ClassName = "none_team"
			v.Acquisition = "N"
		}
		log.Println(&v)
	}

	total := len(list)

	barA := &dto.Team{
		Name:      "A",
		ClassName: "a_team",
		Rate:      (aCnt * 100) / total,
		Uid:       "",
	}

	barB := &dto.Team{
		Name:      "B",
		ClassName: "b_team",
		Rate:      (bCnt * 100) / total,
		Uid:       "",
	}

	barC := &dto.Team{
		Name:      "C",
		ClassName: "c_team",
		Rate:      (cCnt * 100) / total,
		Uid:       "",
	}

	barNone := &dto.Team{
		Name:      "None",
		ClassName: "none_team",
		Rate:      ((total - aCnt - bCnt - cCnt) * 100) / total,
		Uid:       "",
	}

	barList := []*dto.Team{barA, barB, barC, barNone}

	log.Println(barA, barB, barC)

	return map[string]interface{}{
		"JSON":        "Data",
		"barList":     barList,
		"tenantoList": list,
	}, nil
}

func getListJson(tenanto string) ([]*dto.TenantoView, error) {
	bytes, err := useIoutilReadFile(tenanto)
	if err != nil {
		return nil, err
	}
	var tenantoList []*dto.TenantoView
	err = json.Unmarshal(bytes, &tenantoList)
	if err != nil {
		return nil, err
	}
	return tenantoList, nil
}

func getList() []*dto.Tenanto {

	auth := "1zBUevXg2gtvtgCaqw3n591l0PR949ZzgwXabjKm"
	endpoint := "https://mac-001-1e4f9.firebaseio.com/"

	c := firebase.NewClient(endpoint+"/billdiing/e_1", auth, nil)

	retTenanto := func() interface{} {
		return &dto.Tenanto{}
	}

	var list []*dto.Tenanto
	for n := range c.Iterator(retTenanto) {
		dto := n.Value.(*dto.Tenanto)
		log.Println(dto)
		list = append(list, dto)
	}

	return list
}

func useIoutilReadFile(tenanto string) ([]byte, error) {
	bytes, err := ioutil.ReadFile("tool/json/" + tenanto + ".json")
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
