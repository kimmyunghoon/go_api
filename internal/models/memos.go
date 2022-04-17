package models

type Memo struct {
	Title    string `firestore:"title"`
	Contents string `firestore:"contents"`
}
