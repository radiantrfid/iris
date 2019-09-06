package main

import (
	"github.com/radiantrfid/iris/"
	"github.com/radiantrfid/iris//core/netutil"
)

func main() {
	app := iris.New()

	l, err := netutil.UNIX("/tmpl/srv.sock", 0666) // see its code to see how you can manually create a new file listener, it's easy.
	if err != nil {
		panic(err)
	}

	app.Run(iris.Listener(l))
}

// Look "custom-listener/unix-reuseport" too.
