package main

import (
	"Scoreapi_go/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/:type", postItem)

	router.Run("localhost:8080")
}

func postItem(c *gin.Context) {
	itemType, err := items.NewTypeItems(c.Param("type"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, itemType)
	}

	posted := items.Item{}

	if err := c.BindJSON(&posted); err != nil {
		return
	}

	item, err := items.NewItem(posted.Name, posted.Score, posted.Image, posted.Description)
	if err != nil {
		// SHOULD CHECK HOW TO RETURN ERR TEXT
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err})
	}

	itemType.AddItem(*item)
	c.IndentedJSON(http.StatusCreated, item)
}
