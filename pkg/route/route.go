package route

import (
	"context"
	"fmt"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracore/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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

	//srv := asynq.NewServer(
	//	asynq.RedisClientOpt{Addr: utils.RedisAddr},
	//	asynq.Config{
	//		// Specify how many concurrent workers to use
	//		Concurrency: 10,
	//	},
	//)
	//
	//mux := asynq.NewServeMux()

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
	//cateHandlersQueue := yasuoHandler.NewCategoryQueueHandlers(cateService)
	///=================== MssBroker ===================///
	//mux.HandleFunc(utils.JsonCateDelivery, cateHandlersQueue.HandleJsonCateDeliveryTask)
	//go func() {
	//	if err := srv.Run(mux); err != nil {
	//		log.Fatalf("could not run server: %v", err)
	//	}
	//}()

	//=================== HTTPAPI ===================//
	v1Api := s.Router.Group("/api/v1")

	// Add the new API endpoint for saving categories
	v1Api.POST("/shopee/category/save", ginext.WrapHandler(cateHandler.SaveCate))
	v1Api.POST("/shopee/shops/save", ginext.WrapHandler(cateHandler.SaveShopsByCate))
	v1Api.POST("/shopee/shops-detail/save", ginext.WrapHandler(cateHandler.SaveShopsDetail))

	return s
}
