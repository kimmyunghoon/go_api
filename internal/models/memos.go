package models

type Memo struct {
	Id       string `firestore:"id"`
	Title    string `firestore:"title"`
	Contents string `firestore:"contents"`
}
