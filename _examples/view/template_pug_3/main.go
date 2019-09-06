package main

import "github.com/radiantrfid/iris/"

func main() {
	app := iris.New()

	tmpl := iris.Pug("./templates", ".pug")

	app.RegisterView(tmpl)

	app.Get("/", index)

	// http://localhost:8080
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	ctx.View("index.pug")
}
