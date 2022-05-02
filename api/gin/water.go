package gin

import (
	"context"
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

	ctx := context.Background()
	doc, _, err := client.Collection(collection).Add(ctx, tmpWater)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	//doc.DataTo(&returnValues)

	c.JSON(http.StatusOK, gin.H{
		"id":   doc.ID,
		"code": http.StatusOK,
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

	doc, err := client.Collection(collection).Doc(tmpWater.Date).Get(ctx)
	returnValues := models.Water{}
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": returnValues,
		})
	}

	doc.DataTo(&returnValues)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": returnValues,
	})
}
