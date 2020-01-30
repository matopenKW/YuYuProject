package dao

import (
	"YuYuProject/internal/dto"
	"YuYuProject/pkg/util"
	"encoding/json"
	"errors"
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
	// whiteswan に実装してもらいたいとこ
	return nil, nil
}

func GetTeamRateDao() func() ([]*dto.TeamRate, error) {

	c, _ := ini.Load(CONFIG_PATH)
	environment := c.Section("db").Key("environment").MustInt()

	switch environment {
	case 1:
		return getTeamRateJson
	case 2:
		return getTeamRateFireBase
	default:
		return func() ([]*dto.TeamRate, error) {
			return nil, errors.New("daoのセットに失敗しました environment:" + string(environment))
		}
	}
}

func getTeamRateJson() ([]*dto.TeamRate, error) {
	bytes, err := util.ReadFile("tool/json/teamRate.json")
	if err != nil {
		return nil, err
	}

	var teamRateList []*dto.TeamRate
	err = json.Unmarshal(bytes, &teamRateList)
	if err != nil {
		return nil, err
	}

	return teamRateList, nil
}

func getTeamRateFireBase() ([]*dto.TeamRate, error) {
	// whiteswan に実装してもらいたいとこ
	return nil, nil
}
