package counter

import (
	"time"

	"github.com/radiantrfid/iris/"
	"github.com/radiantrfid/iris/core/host"
)

func Configurator(app *iris.Application) {
	counterValue := 0

	go func() {
		ticker := time.NewTicker(time.Second)

		for range ticker.C {
			counterValue++
		}

		app.ConfigureHost(func(h *host.Supervisor) { // <- HERE: IMPORTANT
			h.RegisterOnShutdown(func() {
				ticker.Stop()
			})
		}) // or put the ticker outside of the gofunc and put the configurator before or after the app.Get, outside of this gofunc
	}()

	app.Get("/counter", func(ctx iris.Context) {
		ctx.Writef("Counter value = %d", counterValue)
	})
}
