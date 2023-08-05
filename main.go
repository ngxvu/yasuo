package main

import (
	"context"
	"os"

	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/yasuo/conf"
	"gitlab.com/merakilab9/yasuo/pkg/route"
	"gitlab.com/merakilab9/yasuo/pkg/utils"
)

const (
	APPNAME = "Yasou"
)

func main() {
	conf.SetEnv()
	logger.Init(APPNAME)
	utils.LoadMessageError()

	// Pass the MongoDB database instance to the NewService function
	app := route.NewService()
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		logger.Tag("main").Error(err)
	}
	os.Clearenv()
}
