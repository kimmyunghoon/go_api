package gin

import (
	"github.com/gin-gonic/gin"
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
