package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/pkg/util"
	"log"
	"net/http"

	"firebase.google.com/go/auth"
)

func RegistSerial(userRec *auth.UserRecord, req *http.Request) error {

	req.ParseForm()

	getSingleTeamDao := dao.GetSingleTeamDao()
	team, err := getSingleTeamDao(userRec.UID)
	if err != nil {
		return err
	}

	if err := util.CheckNil(req.Form["serialCode"], "シリアルコード"); err != nil {
		return err
	}

	serialCode := req.Form["serialCode"][0]
	log.Println(serialCode)

	err = regist(team.Id, serialCode)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func regist(teamId, serialCode string) error {
	singleSerialDao := dao.GetSingleSerialDao()
	serial, err := singleSerialDao(serialCode)
	if err != nil {
		return err
	}

	serial.Acquired = true
	serial.Acquisition = teamId
	updateSerialDao := dao.UpdateSerialDao()
	err = updateSerialDao(serialCode, serial)
	if err != nil {
		return err
	}

	err = UpdateTenant(serial)
	if err != nil {
		return err
	}

	return nil
}
