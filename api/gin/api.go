package gin

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api/api/driver"
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
