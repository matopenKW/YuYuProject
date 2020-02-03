package apps

import (
	"YuYuProject/internal/dao"
	"errors"
	"log"
)

func RegistSerial() error {

	// TODO チームIDはキャッシュ、シリアルコードは画面から取得する
	teamId := "A"
	serialCode := "aaaa"

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

	if serial.Acquired {
		return errors.New("すでに登録済みのシリアルコードです。")
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
