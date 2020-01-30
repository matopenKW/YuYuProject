package main

import (
	"YuYuProject/internal/dto"
	"encoding/json"
	"github.com/JustinTulloss/firebase"
	"io/ioutil"
	"log"
	"strconv"
)

const tenanto = "e_1"

func main() {
	auth := "1zBUevXg2gtvtgCaqw3n591l0PR949ZzgwXabjKm"
	endpoint := "https://mac-001-1e4f9.firebaseio.com/"

	c := firebase.NewClient(endpoint, auth, nil)

	log.Println(auth, endpoint)

	bytes := useIoutilReadFile()
	var tenantoList []*dto.Tenanto
	json.Unmarshal(bytes, &tenantoList)

	SetData(c, tenantoList)

}

func Push(c firebase.Client, tenantoList []*dto.Tenanto) {

}

func SetData(c firebase.Client, tenantoList []*dto.Tenanto) {
	var insertMap = map[string]dto.Tenanto{}
	for _, value := range tenantoList {
		insertMap[strconv.Itoa(value.Seq)] = *value
	}

	log.Println(insertMap)

	_, err := c.Set("building/"+tenanto, insertMap, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func useIoutilReadFile() []byte {
	bytes, err := ioutil.ReadFile("json/" + tenanto + ".json")
	if err != nil {
		panic(err)
	}

	return bytes
}
