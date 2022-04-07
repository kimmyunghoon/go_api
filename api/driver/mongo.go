package driver

import (
	"context"
	"fmt"
	"sync"
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	client *mongo.Client
	url    string
}

var dbInstance *MongoDB
var once sync.Once

/**
Singleton pattern 으로 구현한 객체를 가져오는 메소드
참고 : https://pkg.go.dev/sync#Once
*/

func DBClient() *mongo.Client {
	once.Do(func() {
		fmt.Println("DB 연결 시작")
		dbInstance = new(MongoDB)
		newDBConfig(dbInstance)
		newConnectionMongoDBClient(dbInstance)
		fmt.Println("DB 연결 완료")
	})
	return dbInstance.client
}

func newDBConfig(db *MongoDB) {
	db.url = "mongodb://localhost:27017"
}

func newConnectionMongoDBClient(db *MongoDB) {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(db.url))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client 생성")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx) //차후에 데이터베이스 옵션 추가
	if err != nil {
		log.Fatal(err)
	}
	db.client = client
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//err = client.Connect(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	///*
	//   List databases
	//*/
	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(databases)
	//// Interact with data
	//type User struct {
	//	Name string `bson:"name,omitempty"`
	//	Age  int    `bson:"age,omitempty"`
	//}
	//
	///*
	//	Get my collection instance
	//*/
	//collection := client.Database("clone").Collection("users")
	//
	///*
	//	Insert documents
	//*/
	//docs := []interface{}{
	//	bson.D{{"name", "a"}, {"age", 1}},
	//	bson.D{{"name", "b"}, {"age", 2}},
	//	bson.D{{"name", "c"}, {"age", 3}},
	//}
	//
	//res, insertErr := collection.InsertMany(ctx, docs)
	//if insertErr != nil {
	//	log.Fatal(insertErr)
	//}
	//fmt.Println(res)
	///*
	//	Iterate a cursor
	//*/
	//cur, currErr := collection.Find(ctx, bson.D{})
	//
	//if currErr != nil {
	//	panic(currErr)
	//}
	//defer cur.Close(ctx)
	//
	//var posts []User
	//if err = cur.All(ctx, &posts); err != nil {
	//	panic(err)
	//}
	//fmt.Println(posts)
}
