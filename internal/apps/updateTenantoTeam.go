package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"errors"
	"log"
)

func UpdateTenantoTeam() error {
	err := update()
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func update() error {
	serialDao := dao.GetSerialDao()
	serials, err := serialDao()
	if err != nil {
		return err
	}

	for _, serial := range serials {
		err := UpdateTenant(serial)

		if err != nil {
			return nil
		}
	}
	return nil
}

func UpdateTenant(serial *dto.Serial) error {
	errMsg := "Tenantoのレコードが不足しています FloorId:" + string(serial.FloorId) + " Seq:" + string(serial.Seq)

	if !serial.Acquired {
		return nil
	}

	tenantDao := dao.GetTenatoDao()
	tenants, err := tenantDao(serial.FloorId)
	if err != nil {
		return err
	}
	if len(tenants) < serial.Seq {
		return errors.New(errMsg)
	}
	tenant := tenants[serial.Seq-1]
	if tenant.Seq != serial.Seq {
		return errors.New(errMsg)
	}

	tenant.Acquisition = serial.Acquisition
	updateTenanto := dao.UpdateTenantoDao()
	err = updateTenanto(serial.FloorId, tenant)
	if err != nil {
		return err
	}

	return nil
}
