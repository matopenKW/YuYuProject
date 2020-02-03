package dao

import (
	"cloud.google.com/go/firestore"
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"YuYuProject/pkg/db"
	"encoding/json"
	"errors"
	"gopkg.in/ini.v1"
	"github.com/mitchellh/mapstructure"
)

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
			return nil, errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getTenatoLocal(floorId string) ([]*dto.Tenanto, error) {
	bytes, err := util.ReadFile("tool/json/" + floorId + ".json")
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
	collection := func (client *firestore.Client)(*firestore.CollectionRef){
		return client.Collection("building").Doc("twins").Collection(floorId)
	}
	orderBy := func ()(string, firestore.Direction){
		return "Seq", firestore.Asc
	}
	tenantoMaps, err := db.SelectDocuments(client, collection, orderBy)
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
