package driver

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"go_api/api/models"
	"google.golang.org/api/iterator"
	"log"
	"os"
	"testing"
)

func FirebaseDatabase(t *testing.T) {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "C:\\keys\\golang-5bc81-firebase-adminsdk-6a33r-3e2f422ed2.json")
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://golang-5bc81.firebaseio.com",
	}
	//opt := option.WithCredentialsFile("golang-5bc81-firebase-adminsdk-6a33r-3e2f422ed2.json")
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("error get client : %v\n", err)
	}
	t.Log(client)
	ref := client.NewRef("/databases/test/")
	t.Log(ref.Key, ref.Path)
	usersRef := ref.Child("/users")
	err = usersRef.Set(ctx, map[string]*models.User{
		"a": {
			Age:  30,
			Name: "Alan Turing",
		},
		"b": {
			Age:  222,
			Name: "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	//fmt.Println(client.NewRef("test").Get(ctx, j{}))
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
