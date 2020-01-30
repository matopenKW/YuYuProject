package main

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"log"
)

func main() {

	// tenantoList, err := util.GetList("json/e_1", []map[string]interface{}{})
	var list = make([]*dto.Tenanto, 0, 0)
	tenantoList, err := util.GetList("json/e_1", list)

	if err != nil {
		log.Println(err)
	}

	for _, v := range tenantoList {
		log.Println(v)
	}

}
