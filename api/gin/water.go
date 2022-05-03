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

func SetWaterRecord(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := "waters"
	var tmpWater models.Water

	if err := c.BindJSON(&tmpWater); err != nil {
		// DO SOMETHING WITH THE ERROR
	}
	fmt.Println(tmpWater)
	ctx := context.Background()
	_, err := client.Collection(collection).Doc(tmpWater.Date).Set(ctx, tmpWater)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	//doc.DataTo(&returnValues)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": tmpWater,
	})
}

func GetWaterRecord(c *gin.Context) {

	client := driver.FireStoreClient()
	collection := "waters"
	var tmpWater models.Water

	if err := c.BindJSON(&tmpWater); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	ctx := context.Background()
	docinfo := client.Collection(collection).Doc(tmpWater.Date)

	doc, _ := docinfo.Get(ctx)
	returnValues := models.Water{}
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code":    http.StatusBadRequest,
	//		"message": returnValues,
	//	})
	//	return
	//}

	if !doc.Exists() {
		_, _ = client.Collection(collection).Doc(tmpWater.Date).Set(ctx, tmpWater)
	} else {
		doc.DataTo(&returnValues)
	}
	fmt.Println(returnValues)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": returnValues,
	})
	return
}
