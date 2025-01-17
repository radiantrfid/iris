package main

import "github.com/radiantrfid/iris"

func main() {
	app := iris.New()

	// with default template funcs:
	//
	// - {{ urlpath "mynamedroute" "pathParameter_ifneeded" }}
	// - {{ render "header.html" }}
	// - {{ render_r "header.html" }} // partial relative path to current page
	// - {{ yield }}
	// - {{ current }}
	app.RegisterView(iris.HTML("./templates", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Name", "iris") // the .Name inside the ./templates/hi.html
		ctx.Gzip(true)               // enable gzip for big files
		ctx.View("hi.html")          // render the template with the file name relative to the './templates'
	})

	// http://localhost:8080/
	app.Run(iris.Addr(":8080"))
}

/*
Note:

In case you're wondering, the code behind the view engines derives from the "github.com/radiantrfid/iris/view" package,
access to the engines' variables can be granded by "github.com/radiantrfid/iris" package too.

    iris.HTML(...) is a shortcut of view.HTML(...)
    iris.Django(...)     >> >>      view.Django(...)
    iris.Pug(...)        >> >>      view.Pug(...)
    iris.Handlebars(...) >> >>      view.Handlebars(...)
    iris.Amber(...)      >> >>      view.Amber(...)
*/
