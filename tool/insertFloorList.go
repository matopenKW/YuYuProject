package main

import (
	"YuYuProject/internal/dao"
	"YuYuProject/pkg/db"
	"context"
	"log"
	"strconv"
)

func insertFloorList() {

	log.Println("start")

	firestore, err := db.OpenFirestore()
	if err != nil {
		log.Fatalln("connection failed for CloudFireStore , error: %v", err)
	}

	log.Println(firestore)

	tenantoDao := dao.GetTenatoDao()
	floorList, err := tenantoDao("e_1")

	for i, v := range floorList {
		log.Println(v)

		ctx := context.Background()
		ref := firestore.Collection("test_building").Doc("twins").Collection("e_1")
		doc := ref.Doc(strconv.Itoa(i + 1))

		_, err := doc.Set(ctx, v)
		if err != nil {
			log.Fatalln("err")
		}
	}

	log.Println("success")

}
