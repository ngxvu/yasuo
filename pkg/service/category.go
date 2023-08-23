package service

import (
	"context"
	"gitlab.com/merakilab9/yasuo/pkg/model"
	"gitlab.com/merakilab9/yasuo/pkg/repo/mongodb"
)

type CategoryService struct {
	cateRepo mongodb.CategoryRepoInterface
}

type CategoryInterface interface {
	SaveCate(ctx context.Context, cats []model.CategoryList) error
	SaveShopsByCate(ctx context.Context, data model.DataShopCrawled) error
	SaveShopsInfo(ctx context.Context, data model.DataInfoShopsCrawled) error
	SaveShopsInfoByQueue(ctx context.Context, data model.DataInfoShop) error
}

func NewCategoryService(cateRepo mongodb.CategoryRepoInterface) CategoryInterface {
	return &CategoryService{
		cateRepo: cateRepo,
	}
}

func (s *CategoryService) SaveCate(ctx context.Context, cats []model.CategoryList) error {
	if err := s.cateRepo.InsertAllCategories(ctx, cats); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) SaveShopsByCate(ctx context.Context, data model.DataShopCrawled) error {
	if err := s.cateRepo.InsertAllShopsByCate(ctx, data); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) SaveShopsInfo(ctx context.Context, data model.DataInfoShopsCrawled) error {
	if err := s.cateRepo.InsertAllShopsInfo(ctx, data); err != nil {
		return err
	}
	return nil
}

func (s *CategoryService) SaveShopsInfoByQueue(ctx context.Context, data model.DataInfoShop) error {
	if err := s.cateRepo.InsertAllShopsInfoByQueue(ctx, data); err != nil {
		return err
	}
	return nil
}
