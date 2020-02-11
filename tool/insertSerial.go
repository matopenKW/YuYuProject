package tool

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"YuYuProject/pkg/util"
	"context"
	"encoding/json"
	"errors"
	"log"
)

func InsertSerial() error {

	log.Println("start")

	firestore, err := db.OpenFirestore()
	if err != nil {
		return errors.New("connection failed for CloudFireStore , error:" + err.Error())
	}
	log.Println(firestore)

	serialMap, err := getSerial()
	if err != nil {
		return err
	}
	for k, v := range serialMap {
		log.Println(v)

		ctx := context.Background()
		doc := firestore.Collection("serial").Doc(k)

		_, err := doc.Set(ctx, v)
		if err != nil {
			return err
		}
	}

	log.Println("success")

	return nil
}

func getSerial() (map[string]*dto.Serial, error) {
	bytes, err := util.ReadFile(dao.JSON_FOLDER_PATH + "serial.json")
	if err != nil {
		return nil, err
	}

	var serialList []map[string]interface{}
	err = json.Unmarshal(bytes, &serialList)
	if err != nil {
		return nil, err
	}

	serialMap := make(map[string]*dto.Serial)

	for _, v := range serialList {
		var dto dto.Serial
		err := util.MapToStruct(v, &dto)
		if err != nil {
			return nil, err
		}

		serialCode := v["SerialCode"].(string)
		serialMap[serialCode] = &dto
	}

	return serialMap, nil
}
