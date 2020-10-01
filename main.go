package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"order-go/cms/controller"
	"order-go/datasource"
	"order-go/repository"
	"order-go/service"
)

func main() {
	app := newApp()
	mvcHandle(app)
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
	tmpl := iris.HTML("./cms/template", ".html")
	tmpl.Layout(".common/main_layout.html")
	app.HandleDir("/static", "./cms/static")
	app.RegisterView(tmpl)
	app.Logger().SetLevel(Mode)
	return app
}

func mvcHandle(app *iris.Application) {
	db := datasource.NewMysqlEngine()

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	users := mvc.New(app.Party("/user").Layout("common/layout_user.html"))
	users.Register(userService)
	users.Handle(&controller.UserController{})
}
