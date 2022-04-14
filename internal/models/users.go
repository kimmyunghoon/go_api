package models

type User struct {
	Name string `bson:"name,omitempty"`
	Age  int    `bson:"age,omitempty"`
}
