package main

import (
	"github.com/radiantrfid/iris/_benchmarks/iris-mvc-templates/controllers"

	"github.com/radiantrfid/iris"
	"github.com/radiantrfid/iris/context"
	"github.com/radiantrfid/iris/mvc"
)

const (
	// publicDir is the exactly the same path that .NET Core is using for its templates,
	// in order to reduce the size in the repository.
	// Change the "C\\mygopath" to your own GOPATH.
	// publicDir = "C:\\mygopath\\src\\github.com\\kataras\\iris\\_benchmarks\\netcore-mvc-templates\\wwwroot"
	publicDir = "/home/kataras/mygopath/src/github.com/radiantrfid/iris/_benchmarks/netcore-mvc-templates/wwwroot"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/layout.html"))
	app.HandleDir("/public", publicDir)
	app.OnAnyErrorCode(onError)

	mvc.New(app).Handle(new(controllers.HomeController))

	app.Run(iris.Addr(":5000"))
}

type err struct {
	Title string
	Code  int
}

func onError(ctx context.Context) {
	ctx.ViewData("", err{"Error", ctx.GetStatusCode()})
	ctx.View("shared/error.html")
}
