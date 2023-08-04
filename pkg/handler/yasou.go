package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/merakilab9/meradia/pkg/model"
	"io/ioutil"
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
	// Get the raw body of the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Initialize the categoriesData slice
	categoriesData := []model.DataCatCrawled{}

	// Convert the body to a DataCatCrawled object
	dataCatCrawled := model.DataCatCrawled{}
	err = json.Unmarshal(body, &dataCatCrawled)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Append the DataCatCrawled object to the categoriesData slice
	categoriesData = append(categoriesData, dataCatCrawled)

	// Create the categories collection if it doesn't exist
	collection := h.db.Collection("categories")

	// Insert the data into the categories collection
	var dataToInsert []interface{}
	for _, category := range categoriesData {
		dataToInsert = append(dataToInsert, category)
	}
	_, err = collection.InsertMany(context.Background(), dataToInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save categories"})
		return
	}
}
