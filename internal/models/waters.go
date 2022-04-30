package models

type Water struct {
	Value int    `bson:"value,omitempty"`
	Date  string `bson:"date,omitempty"`
}
