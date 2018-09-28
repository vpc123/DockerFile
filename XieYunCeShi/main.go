// file: main.go

package main

import (
	// "github.com/kataras/iris/_examples/mvc/overview/datasource"
	// "github.com/kataras/iris/_examples/mvc/overview/repositories"
	// "github.com/kataras/iris/_examples/mvc/overview/services"
	"./web/controllers"
	// "github.com/kataras/iris/_examples/mvc/overview/web/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	//
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Load the template files.
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Serve our controllers.
	mvc.New(app.Party("/index")).Handle(new(controllers.IndexController))

	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr("localhost:8080"),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		// iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}
