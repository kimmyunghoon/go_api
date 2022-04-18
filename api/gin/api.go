package gin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/api/driver"
	"go_api/internal/models"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

func InsertTest(c *gin.Context) {
	c.String(http.StatusOK, "Insert Test Func")
}

func DeleteTest(c *gin.Context) {
	c.String(http.StatusOK, "Delete Test Func")
}

func UpdateTest(c *gin.Context) {
	c.String(http.StatusOK, "Update Test Func")
}

func CreateTest(c *gin.Context) {
	c.String(http.StatusOK, "Create Test Func")
}

func FirestoneCollectionSet(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	var tmpMemo models.Memo

	if err := c.BindJSON(&tmpMemo); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	ctx := context.Background()
	_, _, err := client.Collection(collection).Add(ctx, tmpMemo)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"title":    tmpMemo.Title,
		"contents": tmpMemo.Contents,
		"code":     http.StatusOK,
	})
}

func GetFirestoneCollectionAllData(c *gin.Context) {
	client := driver.FireStoreClient()
	collection := c.Param("collection")
	ctx := context.Background()
	iter := client.Collection(collection).Documents(ctx)
	fmt.Println(iter)
	data := map[int]string{}
	index := 0
	for {
		doc, err := iter.Next()
		fmt.Println(doc)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data())
		b, _ := json.Marshal(doc.Data())
		data[index] = string(b)
		//data = doc.Data()
		index++
	}
	c.JSON(http.StatusOK, gin.H{
		"find collection": collection,
		"data":            data,
	})
}
func FirestoneCollectionData(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	ctx := context.Background()

	colRef := client.Collection(collection)

	docs, _ := colRef.Documents(ctx).GetAll()
	returnValues := make([]models.Memo, len(docs))
	for index, doc := range docs {
		doc.DataTo(&returnValues[index])
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": returnValues,
	})
}
