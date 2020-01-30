package main

import (
	"YuYuProject/internal/dto"
	"github.com/JustinTulloss/firebase"
	"log"
)

func teamAlloc() interface{} {
	return &dto.Team{}
}

func tenantoAlloc() interface{} {
	return &dto.Tenanto{}
}

func main() {

	auth := "1zBUevXg2gtvtgCaqw3n591l0PR949ZzgwXabjKm"
	endpoint := "https://mac-001-1e4f9.firebaseio.com/building/e"

	// c := firebase.NewClient(endpoint+"/team", auth, nil)
	// for n := range c.Iterator(teamAlloc) {
	// 	log.Printf("FirstName-LastName: %s - %s", n.Value.(*dto.Team).Name, n.Value.(*dto.Team).ClassName)
	// }

	c := firebase.NewClient(endpoint, auth, nil)
	for n := range c.Iterator(tenantoAlloc) {
		log.Println(n)
	}

}
