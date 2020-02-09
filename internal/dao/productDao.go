package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"cloud.google.com/go/firestore"
	"errors"
	"gopkg.in/ini.v1"
	"log"
)

func GetProductDao() func(string, string, *dto.Product) error {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return func(string, string, *dto.Product) error {
			return nil
		}
	case 2:
		return regisatProductFireBase
	default:
		return func(string, string, *dto.Product) error {
			return errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func regisatProductFireBase(teamId, tenantId string, product *dto.Product) error {
	client, err := db.OpenFirestore()
	if err != nil {
		return err
	}

	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection("product").Doc(tenantId).Collection(teamId)
	}

	_, err = db.InsertDocument(client, collection, product)

	log.Println(err)

	return err
}
