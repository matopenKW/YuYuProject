package main

import (
	"YuYuProject/internal/dto"
	"context"
	"encoding/json"
	"github.com/JustinTulloss/firebase"
	"io/ioutil"
	"log"
)

const tenanto = "e"

func main() {
	auth := "1zBUevXg2gtvtgCaqw3n591l0PR949ZzgwXabjKm"
	endpoint := "https://mac-001-1e4f9.firebaseio.com/"

	c := firebase.NewClient(endpoint, auth, nil)

	bytes := useIoutilReadFile()
	var tenantoList []*dto.Tenanto
	json.Unmarshal(bytes, &tenantoList)

	ref := c.NewRef("building/" + tenanto)
	for _, value := range tenantoList {
		childRef1, err := ref.Push(context.Background(), value)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(childRef1)
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
