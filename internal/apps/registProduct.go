package apps

import (
	"YuYuProject/internal/dao"
	"YuYuProject/internal/dto"
	"errors"
	"log"
	"net/http"
	"time"
)

func RegistProduct(req *http.Request) error {

	req.ParseForm()

	// TODO チームIDはキャッシュ、シリアルコードは画面から取得する
	teamId := "C"

	if err := checkNil(req.Form["tenantId"], "テナントId"); err != nil {
		return err
	}

	if err := checkNil(req.Form["productName"], "商品名"); err != nil {
		return err
	}

	if err := checkNil(req.Form["productNo"], "商品番号"); err != nil {
		return err
	}

	tenantId := req.Form["tenantId"][0]
	productName := req.Form["productName"][0]
	productNo := req.Form["productNo"][0]

	err := registProduct(teamId, tenantId, productName, productNo)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func registProduct(teamId, tenantId, productName, productNo string) error {

	now := time.Now()
	product := &dto.Product{
		productName, productNo, &now,
	}
	log.Println(product)

	productDao := dao.GetProductDao()
	err := productDao(teamId, tenantId, product)
	if err != nil {
		return err
	}

	return nil
}

func checkNil(strs []string, errObjName string) error {
	if strs == nil || strs[0] == "" {
		return errors.New(errObjName + "が空です！")
	}

	return nil
}
