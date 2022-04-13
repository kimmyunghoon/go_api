package driver

import (
	"context"
	firebase "firebase.google.com/go"
	"go_api/api/models"
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
