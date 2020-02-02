package apps

import (
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
	// whiteswanに作ってもらうところ

	// 画面で入力したシリアルコード = DocumentIdです。

	// 1. シリアルコレクションから対象のDocumentを取得する

	// 2. Doumentが取得できない場合は以下のエラーをreturn する
	// errors.New("存在しないシリアルコードです。")

	// 3. 取得したDocumentのAcquiredがtrueだった場合は以下のエラーをreturnする
	// errors.New("すでに登録済みのシリアルコードです。")

	// 4. Douemntが存在した場合は対象のDocumentを以下の様に更新
	// Acquired = true
	// Acquisition = teamId

	// 5. その他、エラーが起きた場合はエラーを　returnすること

	return nil
}
