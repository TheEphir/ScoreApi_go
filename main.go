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
	router.GET("/:type/items", showItems)
	router.GET("/:type/:item", lookForItem)
	router.POST("/item", postItem)

	router.Run("localhost:8080")
}

func postItem(c *gin.Context) {
	posted := postedItem{}
	if err := c.BindJSON(&posted); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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

func showItems(c *gin.Context) {
	itemType := c.Param("type")

	res, err := items.NewTypeItems(itemType)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.IndentedJSON(http.StatusOK, res.Items)
}

func lookForItem(c *gin.Context) {
	itemType := c.Param("type")
	itemName := c.Param("item")

	db, err := items.ReadType(itemType)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "something wrong with type"})
	}

	res := db.SearchItemByName(itemName)
	if len(res) < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.IndentedJSON(http.StatusOK, res)
}
