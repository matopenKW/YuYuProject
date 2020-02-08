package apps

import (
	"YuYuProject/internal/dao"
	"errors"
	"log"
	"net/http"
)

func RegistSerial(req *http.Request) error {

	req.ParseForm()

	// TODO チームIDはキャッシュ、シリアルコードは画面から取得する
	teamId := "C"

	log.Println(req.Form["serialCode"][0])

	if req.Form["serialCode"] == nil || req.Form["serialCode"][0] == "" {
		return errors.New("シリアルコードが空です！")
	}

	serialCode := req.Form["serialCode"][0]
	log.Println(serialCode)

	err := regist(teamId, serialCode)
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

	return nil
}
