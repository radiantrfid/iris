/* Package cache provides server-side caching capabilities with rich support of options and rules.

Use it for server-side caching, see the `iris#Cache304` for an alternative approach that
may fit your needs most.

Example code:


		 import (
		 	"time"

		 	"github.com/radiantrfid/iris"
		 	"github.com/radiantrfid/iris/cache"
		 )

		 func main(){
		 	app := iris.Default()
		 	middleware := cache.Handler(2 *time.Minute)
		 	app.Get("/hello", middleware, h)
		 	app.Run(iris.Addr(":8080"))
		 }

		 func h(ctx iris.Context) {
		 	ctx.HTML("<h1> Hello, this should be cached. Every 2 minutes it will be refreshed, check your browser's inspector</h1>")
		 }
*/

package cache

import (
	"time"

	"github.com/radiantrfid/iris/cache/client"
	"github.com/radiantrfid/iris/context"
)

// Cache accepts the cache expiration duration.
// If the "expiration" input argument is invalid, <=2 seconds,
// then expiration is taken by the "cache-control's maxage" header.
// Returns a Handler structure which you can use to customize cache furher.
//
// All types of response can be cached, templates, json, text, anything.
//
// Use it for server-side caching, see the `iris#Cache304` for an alternative approach that
// may be more suited to your needs.
//
// You can add validators with this function.
func Cache(expiration time.Duration) *client.Handler {
	return client.NewHandler(expiration)
}

// Handler like `Cache` but returns an Iris Handler to be used as a middleware.
// For more options use the `Cache`.
//
// Examples can be found at: https://github.com/radiantrfid/iris/tree/master/_examples/#caching
func Handler(expiration time.Duration) context.Handler {
	h := Cache(expiration).ServeHTTP
	return h
}
