// file: web/controllers/hello_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
)

// HelloController is our sample controller
// it handles GET: /hello and GET: /hello/{name}
type IndexController struct{}

var indexView = mvc.View{
	Name: "index/index.html",
	Data: map[string]interface{}{
		"Title":     "Index",
		"MyMessage": "XieYun",
	},
}

func (c *IndexController) Get() mvc.Result {
	return indexView
}

func (c *IndexController) GetBy(name string) mvc.Result {
	return mvc.View{
		Name: "index/name.html",
		Data: name,
	}
}
