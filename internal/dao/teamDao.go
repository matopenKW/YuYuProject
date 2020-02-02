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

func GetTeamDao() func() ([]*dto.Team, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return getTeamJson
	case 2:
		return getTeamFireBase
	default:
		return func() ([]*dto.Team, error) {
			return nil, errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getTeamJson() ([]*dto.Team, error) {
	bytes, err := util.ReadFile("tool/json/team.json")
	if err != nil {
		return nil, err
	}

	var teamList []*dto.Team
	err = json.Unmarshal(bytes, &teamList)
	if err != nil {
		return nil, err
	}

	return teamList, nil
}

func getTeamFireBase() ([]*dto.Team, error) {
	client, err := db.OpenFirestore()
	collection := func (client *firestore.Client)(*firestore.CollectionRef){
		return client.Collection("team")
	}
	orderBy := func ()(string, firestore.Direction){
		return "Id", firestore.Asc
	}
	teamMaps, err := db.SelectDocuments(client, collection, orderBy)
	if err != nil {
		return nil, err
	}
	var teamList []*dto.Team
	for _, teamMap := range teamMaps {
		var team *dto.Team
		err := mapstructure.Decode(teamMap, &team)
		if err != nil {
			return nil, err
		}
		teamList = append(teamList, team)
	}
	return teamList, nil
}
