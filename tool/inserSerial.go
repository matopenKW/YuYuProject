package main

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"YuYuProject/pkg/util"
	"context"
	"encoding/json"
	"log"
)

func inserSerial() {

	log.Println("start")

	firestore, err := db.OpenFirestore()
	if err != nil {
		log.Fatalln("connection failed for CloudFireStore , error: %v", err)
	}
	log.Println(firestore)

	serialList, err := getSerial()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range serialList {
		log.Println(v)

		ctx := context.Background()
		ref := firestore.Collection("test_serial")
		doc := ref.NewDoc()

		_, err := doc.Set(ctx, v)
		if err != nil {
			log.Fatalln("err")
		}
	}

	log.Println("success")
}

func getSerial() ([]*dto.Serial, error) {
	bytes, err := util.ReadFile("tool/json/serial.json")
	if err != nil {
		return nil, err
	}

	var serialList []*dto.Serial
	err = json.Unmarshal(bytes, &serialList)
	if err != nil {
		return nil, err
	}

	return serialList, nil
}
