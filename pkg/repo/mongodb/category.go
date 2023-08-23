package mongodb

import (
	"context"
	"gitlab.com/merakilab9/yasuo/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CategoryRepo struct {
	DB *mongo.Database
}

type CategoryRepoInterface interface {
	InsertAllCategories(ctx context.Context, categories []model.CategoryList) error
	InsertAllShopsByCate(ctx context.Context, data model.DataShopCrawled) error
	InsertAllShopsInfo(ctx context.Context, data model.DataInfoShopsCrawled) error
	InsertAllShopsInfoByQueue(ctx context.Context, data model.DataInfoShop) error
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

func (r *CategoryRepo) InsertAllShopsByCate(ctx context.Context, data model.DataShopCrawled) error {
	var rs []interface{}
	for _, shops := range data.Data {
		rs = append(rs, shops.Data)
	}
	_, err := r.DB.Collection("shops").InsertMany(ctx, rs)
	if err != nil {
		log.Println("Co loi MongoDB.", err)
		return err
	}
	return nil
}

func (r *CategoryRepo) InsertAllShopsInfo(ctx context.Context, data model.DataInfoShopsCrawled) error {
	var rs []interface{}
	for _, shopinfo := range data.DataInfoShops {
		rs = append(rs, shopinfo)
	}
	_, err := r.DB.Collection("shop-info").InsertMany(ctx, rs)
	if err != nil {
		log.Println("Co loi MongoDB.", err)
		return err
	}
	return nil
}

func (r *CategoryRepo) InsertAllShopsInfoByQueue(ctx context.Context, data model.DataInfoShop) error {
	var rs []interface{}
	rs = append(rs, data)
	_, err := r.DB.Collection("shop-info-by-queue").InsertMany(ctx, rs)
	if err != nil {
		log.Println("Co loi MongoDB.", err)
		return err
	}
	return nil
}
