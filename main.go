package main

import (
	"context"
	"gitlab.com/merakilab9/meradia/conf"
	"gitlab.com/merakilab9/meradia/pkg/repo/mongodb"
	"gitlab.com/merakilab9/meradia/pkg/route"
	"gitlab.com/merakilab9/meradia/pkg/utils"
	"os"

	"gitlab.com/merakilab9/meracore/logger"
)

const (
	APPNAME = "Yasou"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	err := mongodb.InitDatabase() // Initialize MongoDB connection
	if err != nil {
		logger.Tag("main").Error(err)
		return
	}
	defer mongodb.GetDB().Client().Disconnect(context.Background()) // Close MongoDB connection

	// Pass the MongoDB database instance to the NewService function
	app := route.NewService(mongodb.GetDB())
	ctx := context.Background()
	err = app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
