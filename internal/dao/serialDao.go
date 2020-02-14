package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/db"
	"YuYuProject/pkg/util"
	"encoding/json"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/ini.v1"
)

const SERIAL_COLLECTION = "test_serial"

func GetSerialDao() func() ([]*dto.Serial, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return getSerialJson
	case 2:
		return getSerialFireBase
	default:
		return func() ([]*dto.Serial, error) {
			return nil, errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getSerialJson() ([]*dto.Serial, error) {

	return nil, nil
}

func getSerialFireBase() ([]*dto.Serial, error) {
	client, err := db.OpenFirestore()
	if err != nil {
		return nil, err
	}
	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection(SERIAL_COLLECTION)
	}
	option := &db.Option{
		OrderBy: func() (string, firestore.Direction) {
			return "Seq", firestore.Asc
		},
	}

	serialMaps, err := db.SelectDocuments(client, collection, option)
	if err != nil {
		return nil, err
	}
	var serialList []*dto.Serial
	for _, serialMap := range serialMaps {
		var serial *dto.Serial
		err := mapstructure.Decode(serialMap, &serial)
		if err != nil {
			return nil, err
		}
		serialList = append(serialList, serial)
	}
	return serialList, nil
}

func GetSingleSerialDao() func(string) (*dto.Serial, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return getSingleSerialJson
	case 2:
		return getSingleSerialFireBase
	default:
		return func(string) (*dto.Serial, error) {
			return nil, errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getSingleSerialJson(serialCode string) (*dto.Serial, error) {
	bytes, err := util.ReadFile("tool/json/serial.json")
	if err != nil {
		return nil, err
	}

	var serialList []*dto.Serial
	err = json.Unmarshal(bytes, &serialList)
	if err != nil {
		return nil, err
	}

	var serial *dto.Serial
	for _, v := range serialList {
		serial = v
	}
	return serial, nil
}

func getSingleSerialFireBase(serialCode string) (*dto.Serial, error) {
	client, err := db.OpenFirestore()
	if err != nil {
		return nil, err
	}
	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection(SERIAL_COLLECTION)
	}

	serialMap, err := db.SelectDocument(client, serialCode, collection)
	if err != nil {
		return nil, err
	}
	var serial *dto.Serial
	err = mapstructure.Decode(serialMap, &serial)
	if err != nil {
		return nil, err
	}
	return serial, nil
}

func UpdateSerialDao() func(string, *dto.Serial) error {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return func(string, *dto.Serial) error {
			return nil
		}
	case 2:
		return updateSerialFireBase
	default:
		return func(string, *dto.Serial) error {
			return errors.New("daoのセットに失敗しました dao: UpdateSerialDao, environment:" + string(environment))
		}
	}
}

func updateSerialFireBase(serialCode string, serial *dto.Serial) error {
	client, err := db.OpenFirestore()
	if err != nil {
		return err
	}
	document := func(client *firestore.Client) *firestore.DocumentRef {
		return client.Collection(SERIAL_COLLECTION).Doc(serialCode)
	}
	err = db.UpdateDocument(client, document, serial)
	if err != nil {
		return err
	}
	return nil
}
