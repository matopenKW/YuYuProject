package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"encoding/json"
	"errors"
	"gopkg.in/ini.v1"
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
	// whiteswan に実装してもらいたいとこ
	return nil, nil
}
