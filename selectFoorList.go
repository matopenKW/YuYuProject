package main

import (
	"YuYuProject/pkg/db"
	"cloud.google.com/go/firestore"
	"log"
)

func main() {

	fireStore, err := db.OpenFirestore()
	if err != nil {
		log.Fatalln(err)
	}

	collection := func(client *firestore.Client) *firestore.CollectionRef {
		return client.Collection("test_building").Doc("twins").Collection("e_1")
	}

	// Option未指定
	list, err := db.SelectDocuments(fireStore, collection, nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range list {
		log.Println(v)
	}

	// Option指定
	orderBy := func() (string, firestore.Direction) {
		return "Seq", firestore.Asc
	}
	list, err = db.SelectDocuments(fireStore, collection, orderBy)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range list {
		log.Println(v)
	}
}
