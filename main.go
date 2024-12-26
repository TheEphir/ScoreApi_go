package main

import (
	"Scoreapi_go/files"
	"Scoreapi_go/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postedItem struct {
	Type        string `json:"item_type"`
	Name        string `json:"name"`
	Score       int    `json:"score"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func main() {
	router := gin.Default()

	router.GET("/types", showTypes)
	router.POST("/item", postItem)

	router.Run("localhost:8080")
}

func postItem(c *gin.Context) {
	posted := postedItem{}
	if err := c.BindJSON(&posted); err != nil {
		return
	}

	itemType, err := items.NewTypeItems(posted.Type)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	item, err := items.NewItem(posted.Name, posted.Score, posted.Image, posted.Description)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
	}

	itemType.AddItem(*item)
	c.IndentedJSON(http.StatusCreated, item)
}

func showTypes(c *gin.Context) {
	res, err := files.ReadDb()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.IndentedJSON(http.StatusOK, res)
}
