package mongodb

import (
	"context"
	"gitlab.com/merakilab9/yasuo/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepo struct {
	DB *mongo.Database
}

type CategoryRepoInterface interface {
	InsertAllCategories(ctx context.Context, categories []model.CategoryList) error
}

func NewPGRepo(db *mongo.Database) CategoryRepoInterface {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) InsertAllCategories(ctx context.Context, categories []model.CategoryList) error {
	var data []interface{}
	for _, cate := range categories {
		data = append(data, cate)
	}
	_, err := r.DB.Collection("category").InsertMany(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
