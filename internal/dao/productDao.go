package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"errors"
	"log"

	"time"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/ini.v1"
)

const COLLECTION_NAME = "test_product"

func RagistProductDao() func(teamId, tenantId string, product *dto.Product) error {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return func(teamId, tenantId string, product *dto.Product) error {
			return nil
		}
	case 2:
		return regisatProductFireBase
	default:
		return func(teamId, tenantId string, product *dto.Product) error {
			return errors.New("daoのセットに失敗しました dao:RagistProductDao environment:" + string(environment))
		}
	}
}

func regisatProductFireBase(teamId, tenantId string, product *dto.Product) error {
	client, err := db.OpenFirestore()
	if err != nil {
		return err
	}

	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection(COLLECTION_NAME).Doc(tenantId).Collection(teamId)
	}

	_, err = db.InsertDocument(client, collection, product)

	log.Println(err)

	return err
}

func GetProductDao() func(teamId, tenantId string) ([]*dto.Product, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return getProductDaoLocal
	case 2:
		return getProductDaoFirebase
	default:
		return func(teamId, tenantId string) ([]*dto.Product, error) {
			return nil, errors.New("daoのセットに失敗しました dao:GetProductDao environment:" + string(environment))
		}
	}
}

func getProductDaoLocal(teamId, tenantId string) ([]*dto.Product, error) {

	now := time.Now()
	team := &dto.Product{
		"にとにとり", "111222", &now,
	}
	var list []*dto.Product
	for i := 0; i < 15; i++ {
		list = append(list, team)
	}

	return list, nil
}

func getProductDaoFirebase(teamId, tenantId string) ([]*dto.Product, error) {
	client, err := db.OpenFirestore()
	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection(COLLECTION_NAME).Doc(tenantId).Collection(teamId)
	}
	option := &db.Option{
		OrderBy: func() (string, firestore.Direction) {
			return "Timestamp", firestore.Asc
		},
	}
	productMaps, err := db.SelectDocuments(client, collection, option)
	if err != nil {
		return nil, err
	}
	var productList []*dto.Product
	for _, productMap := range productMaps {
		var product *dto.Product
		err := mapstructure.Decode(productMap, &product)
		if err != nil {
			return nil, err
		}
		productList = append(productList, product)
	}
	return productList, nil
}
