package main

import (
	"github.com/kataras/iris/v12"
	"order-go/datasource"
)

func main() {
	app := newApp()
	mvcHandle()
	cfg := iris.YAML("./configs/iris.yml")
	addr := cfg.Other["Addr"].(string)
	app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func newApp() *iris.Application {
	cfg := iris.YAML("./configs/iris.yml")
	Mode := cfg.Other["Mode"].(string)
	app := iris.New()
	app.Logger().SetLevel(Mode)
	return app
}

func mvcHandle() {
	datasource.NewMysqlEngine()
}
