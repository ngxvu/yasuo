package route

import (
	"context"
	"fmt"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/merakilab9/meracore/service"

	"gitlab.com/merakilab9/yasuo/conf"
	yasuoHandler "gitlab.com/merakilab9/yasuo/pkg/handler"
	"gitlab.com/merakilab9/yasuo/pkg/repo/mongodb"
	yasuoService "gitlab.com/merakilab9/yasuo/pkg/service"
)

type Service struct {
	*service.BaseApp
}

// NewService creates a new Service instance with the provided MongoDB database instance.
func NewService() *Service {
	s := &Service{
		BaseApp: service.NewApp("Yasuo Service", "v1.0"),
	}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", conf.LoadEnv().MongoDBHost, conf.LoadEnv().MongoDBPort)),
	)
	if err != nil {
		logger.WithCtx(context.Background(), "connect mongodb failed").Error(err.Error())
	}

	db := client.Database(conf.LoadEnv().MongoDBName)

	cateRepo := mongodb.NewPGRepo(db)
	cateService := yasuoService.NewCategoryService(cateRepo)
	cateHandler := yasuoHandler.NewCategoryHandlers(cateService)

	v1Api := s.Router.Group("/api/v1")

	// Add the new API endpoint for saving categories
	v1Api.POST("/shopee/category/save", ginext.WrapHandler(cateHandler.SaveCate))

	return s
}
