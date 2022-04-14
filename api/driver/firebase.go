package driver

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"testing"
)

type Firestore struct {
	client *firestore.Client
	conf   *firebase.Config
}

var firestoreInstance *Firestore

func FireStoreClient() *firestore.Client {
	once.Do(func() {
		fmt.Println("DB 연결 시작")
		firestoreInstance = new(Firestore)
		newFirestoreConfig(firestoreInstance)
		newConnectionFirestoreClient(firestoreInstance)
		fmt.Println("DB 연결 완료")
	})
	return firestoreInstance.client
}
func newFirestoreConfig(store *Firestore) {
	store.conf = &firebase.Config{
		DatabaseURL: "https://golang-5bc81.firebaseio.com",
	}
}

func newConnectionFirestoreClient(store *Firestore) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\keys\\golang-5bc81-firebase-adminsdk-6a33r-3e2f422ed2.json")

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, store.conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	store.client = client
}

func FirestoreInit(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\keys\\golang-5bc81-firebase-adminsdk-6a33r-3e2f422ed2.json")
	conf := &firebase.Config{
		DatabaseURL: "https://golang-5bc81.firebaseio.com",
	}
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
	}
	dsnap, err := client.Collection("users").Doc("Nb7wlxtVD6eMwRBBw1lN").Get(ctx)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)
	//_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	//	"first": "Ada",
	//	"last":  "Lovelace",
	//	"born":  1815,
	//})
	//if err != nil {
	//	log.Fatalf("Failed adding alovelace: %v", err)
	//}

}
