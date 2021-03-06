package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"YuYuProject/pkg/util"
	"encoding/json"
	"errors"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/ini.v1"
)

const BUILDING_COLLECTION = "building"

func GetTenatoDao() func(floorId string) ([]*dto.Tenanto, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()

	switch environment {
	case 1:
		return getTenatoLocal
	case 2:
		return getTenatoFireBase
	default:
		return func(floorId string) ([]*dto.Tenanto, error) {
			return nil, errors.New("GetTenatoDaoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getTenatoLocal(floorId string) ([]*dto.Tenanto, error) {
	bytes, err := util.ReadFile(JSON_FOLDER_PATH + "twins/" + floorId + ".json")
	if err != nil {
		return nil, err
	}
	var tenantoList []*dto.Tenanto
	err = json.Unmarshal(bytes, &tenantoList)
	if err != nil {
		return nil, err
	}
	return tenantoList, nil
}

func getTenatoFireBase(floorId string) ([]*dto.Tenanto, error) {
	client, err := db.OpenFirestore()
	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection(BUILDING_COLLECTION).Doc("twins").Collection(floorId)
	}
	option := &db.Option{
		OrderBy: func() (string, firestore.Direction) {
			return "Seq", firestore.Asc
		},
	}
	tenantoMaps, err := db.SelectDocuments(client, collection, option)
	if err != nil {
		return nil, err
	}
	var tenantoList []*dto.Tenanto
	for _, tenantoMap := range tenantoMaps {
		var tenanto *dto.Tenanto
		err := mapstructure.Decode(tenantoMap, &tenanto)
		if err != nil {
			return nil, err
		}
		tenantoList = append(tenantoList, tenanto)
	}
	return tenantoList, nil
}

func UpdateTenantoDao() func(string, *dto.Tenanto) error {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return func(string, *dto.Tenanto) error {
			return nil
		}
	case 2:
		return updateTenantoFireBase
	default:
		return func(string, *dto.Tenanto) error {
			return errors.New("UpdateTenantoDaoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func updateTenantoFireBase(floorId string, tenanto *dto.Tenanto) error {
	client, err := db.OpenFirestore()
	if err != nil {
		return err
	}
	document := func(client *firestore.Client) *firestore.DocumentRef {
		return client.Collection(BUILDING_COLLECTION).Doc("twins").Collection(floorId).Doc(strconv.Itoa(tenanto.Seq))
	}
	err = db.UpdateDocument(client, document, tenanto)

	if err != nil {
		return err
	}
	return nil
}
