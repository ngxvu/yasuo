package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/merakilab9/meradia/pkg/model"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type YasouHandler struct {
	db *mongo.Database
}

func NewYasouHandler(db *mongo.Database) *YasouHandler {
	return &YasouHandler{db: db}
}

func (h *YasouHandler) SaveCategoriesHandler(c *gin.Context) {
	// Parse the request body to get the categories data
	var categoriesData model.DataCatCrawled
	err := c.ShouldBindJSON(&categoriesData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Save the categories data to MongoDB
	collection := h.db.Collection("categories")

	for _, category := range categoriesData.Data.CategoryList {
		_, err = collection.InsertOne(c, category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save category"})
			return
		}
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Categories saved successfully"})
}
