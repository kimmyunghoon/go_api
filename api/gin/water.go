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
	var tmpMemo models.Water

	if err := c.BindJSON(&tmpMemo); err != nil {
		// DO SOMETHING WITH THE ERROR
	}

	ctx := context.Background()
	doc, _, err := client.Collection(collection).Add(ctx, tmpMemo)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	//docs, _ := client.Collection(collection).Documents(ctx).Get()
	returnValues := models.Water{}

	//doc.DataTo(&returnValues)

	c.JSON(http.StatusOK, gin.H{
		"id":      doc.ID,
		"code":    http.StatusOK,
		"message": returnValues,
	})
}
