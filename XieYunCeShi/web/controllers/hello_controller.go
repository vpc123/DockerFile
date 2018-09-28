// file: web/controllers/hello_controller.go

package controllers

import (
	"errors"

	"github.com/kataras/iris/mvc"
)

// HelloController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type HelloController struct{}

var helloView = mvc.View{
	Name: "hello/index.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

func (c *HelloController) Get() mvc.Result {
	return helloView
}

// you can define a standard error in order to re-use anywhere in your app.
var errBadName = errors.New("bad name")
var badName = mvc.Response{Err: errBadName, Code: 400}

func (c *HelloController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
	}
	// return mvc.Response{Text: "Hello " + name} OR:
	return mvc.View{
		Name: "hello/name.html",
		Data: name,
	}
}
