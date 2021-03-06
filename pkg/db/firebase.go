package db

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const keyjson = "pkg/conf/key.json"

// 行く行くはリスト項目で持つのがよいのでは
type Option struct {
	Where   func() (fieldName string, condition string, item interface{})
	OrderBy func() (string, firestore.Direction)
}

func getWhere(opt *Option) func() (string, string, interface{}) {
	return opt.Where
}

func getOrderBy(opt *Option) func() (string, firestore.Direction) {
	return opt.OrderBy
}

func OpenAuth() (*auth.Client, error) {
	ctx := context.Background()

	opt := option.WithCredentialsFile(keyjson)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Firestore Auth Succeeded!")
	return client, nil
}

func OpenFirestore() (*firestore.Client, error) {
	opt := option.WithCredentialsFile(keyjson)
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetUserRecord(client *auth.Client, uid string) (*auth.UserRecord, error) {
	user, err := client.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return user, err
}

func UpdateUserInfo(client *auth.Client, uid string, userToUodate *auth.UserToUpdate) (*auth.UserRecord, error) {
	user, err := client.UpdateUser(context.Background(), uid, userToUodate)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func SelectDocument(client *firestore.Client, id string, collection func(client *firestore.Client) *firestore.CollectionRef) (map[string]interface{}, error) {
	ctx := context.Background()

	colle := collection(client)
	if colle == nil {
		return nil, errors.New("failed to connect")
	}

	doc, err := colle.Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.New("存在しないシリアルコードです。 ID:" + string(id))
	}
	data := doc.Data()
	data["DocumentId"] = doc.Ref.ID

	log.Println("Select Document:")
	log.Println(data)
	return data, nil
}

func SelectDocuments(client *firestore.Client, collection func(client *firestore.Client) *firestore.CollectionRef, opt *Option) ([]map[string]interface{}, error) {
	ctx := context.Background()

	list := make([]map[string]interface{}, 0, 10)

	colle := collection(client)
	if colle == nil {
		return nil, errors.New("failed to connect")
	}

	query := colle.Query
	if opt.Where != nil {
		query = query.Where(opt.Where())
	}
	if opt.OrderBy != nil {
		query = query.OrderBy(opt.OrderBy())
	}

	iter := query.Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		data := doc.Data()
		data["DocumentId"] = doc.Ref.ID
		list = append(list, data)
	}

	log.Println("Select Documents:")
	log.Println(list)
	return list, nil
}

func DeleteCollection(client *firestore.Client, ref *firestore.CollectionRef, batchSize int) error {
	ctx := context.Background()

	for {
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}

func DeleteDocument(client *firestore.Client, userId string, documentId string) error {
	ctx := context.Background()

	doc := client.Collection(userId).Doc(documentId)
	batch := client.Batch()
	batch.Delete(doc)

	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDocument(client *firestore.Client, document func(client *firestore.Client) *firestore.DocumentRef, data interface{}) error {
	ctx := context.Background()

	log.Println("Update Document:")
	log.Println(data)
	_, err := document(client).Set(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func InsertDocument(client *firestore.Client, collection func(client *firestore.Client) *firestore.CollectionRef, data interface{}) (string, error) {
	ctx := context.Background()

	log.Println("Insert Document:")
	log.Println(data)

	collec := collection(client)
	if collec == nil {
		return "", errors.New("failed to connect")
	}

	doc, _, err := collec.Add(ctx, data)

	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
