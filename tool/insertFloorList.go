package tool

import (
	"YuYuProject/internal/dao"
	"YuYuProject/pkg/db"
	"context"
	"errors"
	"log"
	"strconv"
)

func InsertFloorList(flootId string) error {

	firestore, err := db.OpenFirestore()
	if err != nil {
		return errors.New("connection failed for CloudFireStore , error: " + err.Error())
	}

	tenantoDao := dao.GetTenatoDao()
	floorList, err := tenantoDao(flootId)

	for i, v := range floorList {
		log.Println(v)

		ctx := context.Background()
		ref := firestore.Collection("building").Doc("twins").Collection(flootId)
		doc := ref.Doc(strconv.Itoa(i + 1))

		_, err := doc.Set(ctx, v)
		if err != nil {
			return err
		}
	}

	return nil

}
