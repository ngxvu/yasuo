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
