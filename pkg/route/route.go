package route

import (
	"github.com/gin-contrib/pprof"
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/merakilab9/meracore/service"
	"gitlab.com/merakilab9/meradia/conf"
	"gitlab.com/merakilab9/meradia/pkg/handler"
)

type Service struct {
	*service.BaseApp
	db *mongo.Database // Add a field to store the MongoDB database instance
}

// NewService creates a new Service instance with the provided MongoDB database instance.
func NewService(db *mongo.Database) *Service {
	s := &Service{
		BaseApp: service.NewApp("Yasou Service", "v1.0"),
		db:      db, // Store the MongoDB database instance
	}

	if !conf.LoadEnv().DbDebugEnable {
		s.db = s.db.Client().Database(s.db.Name())
	}

	// Create a new YasouHandler instance to handle HTTP requests for saving categories
	yasouHandler := handler.NewYasouHandler(s.db)

	pprof.Register(s.Router)

	v1Api := s.Router.Group("/api/v1")

	// Add the new API endpoint for saving categories
	v1Api.POST("/shopee/category/save", yasouHandler.SaveCategoriesHandler)

	return s
}
