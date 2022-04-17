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

func GetFirestoneCollectionSet(c *gin.Context) {
	client := driver.FireStoreClient()
	collection := c.Param("collection")
	title := c.DefaultQuery("title", "none title")
	contents := c.DefaultQuery("contents", "none contents")
	ctx := context.Background()
	_, err := client.Collection(collection).Doc("test").Set(ctx, models.Memo{
		Title:    title,
		Contents: contents,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"title":    title,
		"contents": contents,
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
	//gin.h = map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"find collection": collection,
		"data":            data,
	})
}
func GetFirestoneCollectionData(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := c.Param("collection")
	ctx := context.Background()
	//iter := client.Collection(collection).Documents(ctx).GetAll()

	dsnap, err := client.Collection(collection).Doc("common").Get(ctx)
	if err != nil {
		log.Fatalf("err : %v", err)
	}
	//make([]models.Memo,0)
	memo := models.Memo{}
	dsnap.DataTo(&memo)

	e, err := json.Marshal(memo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))
	fmt.Println("result ", dsnap.Data(), memo)
	//gin.h = map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": string(e), // cast it to string before showing
	})
}
