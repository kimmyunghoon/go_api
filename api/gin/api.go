package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/api/driver"
	"go_api/internal/models"
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

func FirestoneCollectionDelete(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	var tmpMemo models.Memo

	if err := c.BindJSON(&tmpMemo); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	ctx := context.Background()
	fmt.Println(tmpMemo.Id)
	_, err := client.Collection(collection).Doc(tmpMemo.Id).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		c.JSON(http.StatusOK, gin.H{
			"id":       tmpMemo.Id,
			"title":    tmpMemo.Title,
			"contents": tmpMemo.Contents,
			"code":     http.StatusNoContent,
			"message":  nil,
		})
	}
	docs, _ := client.Collection(collection).Documents(ctx).GetAll()
	returnValues := make([]models.Memo, len(docs))
	for index, doc := range docs {
		doc.DataTo(&returnValues[index])
		returnValues[index].Id = doc.Ref.ID
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       tmpMemo.Id,
		"title":    tmpMemo.Title,
		"contents": tmpMemo.Contents,
		"code":     http.StatusOK,
		"message":  returnValues,
	})
}
func FirestoneCollectionSet(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	var tmpMemo models.Memo

	if err := c.BindJSON(&tmpMemo); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	ctx := context.Background()
	doc, _, err := client.Collection(collection).Add(ctx, tmpMemo)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	docs, _ := client.Collection(collection).Documents(ctx).GetAll()
	returnValues := make([]models.Memo, len(docs))
	for index, doc := range docs {
		doc.DataTo(&returnValues[index])
		returnValues[index].Id = doc.Ref.ID
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       doc.ID,
		"title":    tmpMemo.Title,
		"contents": tmpMemo.Contents,
		"code":     http.StatusOK,
		"message":  returnValues,
	})
}
func FirestoneCollectionAll(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	ctx := context.Background()
	docs, _ := client.Collection(collection).Documents(ctx).GetAll()
	returnValues := make([]models.Memo, len(docs))
	for index, doc := range docs {
		doc.DataTo(&returnValues[index])
		returnValues[index].Id = doc.Ref.ID
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": returnValues,
	})
}
