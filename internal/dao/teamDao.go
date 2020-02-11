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
			return nil, errors.New("daoのセットに失敗しました dao: GetTeamDao, environment:" + string(environment))
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
	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection("team")
	}
	option := &db.Option{
		OrderBy: func() (string, firestore.Direction) {
			return "Id", firestore.Asc
		},
	}
	teamMaps, err := db.SelectDocuments(client, collection, option)
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

func GetSingleTeamDao() func(userId string) (*dto.Team, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()
	switch environment {
	case 1:
		return getSingleTeamLocal
	case 2:
		return getSingleTeamFireBase
	default:
		return func(userId string) (*dto.Team, error) {
			return nil, errors.New("daoのセットに失敗しました dao: GetSingleTeamDao, environment:" + string(environment))
		}
	}
}

func getSingleTeamLocal(userId string) (*dto.Team, error) {
	return &dto.Team{
		"C", "c_team", userId, 0, 0, 0,
	}, nil
}

func getSingleTeamFireBase(userId string) (*dto.Team, error) {
	client, err := db.OpenFirestore()
	if err != nil {
		return nil, err
	}

	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection("team")
	}

	option := &db.Option{
		Where: func() (string, string, interface{}) {
			return "Uid", "==", userId
		},
	}

	serialMaps, err := db.SelectDocuments(client, collection, option)
	if err != nil {
		return nil, err
	}

	if len(serialMaps) < 1 || serialMaps[0] == nil {
		return nil, errors.New("Is not found team")
	}

	ret := &dto.Team{}
	err = util.MapToStruct(serialMaps[0], ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
