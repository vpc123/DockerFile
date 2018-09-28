package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/kataras/iris"
)

const maxSize = 5 << 20 // 5MB

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./templates", ".html"))

	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		ctx.View("upload_form.html", token)
	})

	// Handle the post request from the upload_form.html to the server
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		// Get the file from the request.
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename
		var filenameWithSuffix string
		filenameWithSuffix = path.Base(fname) //获取文件名带后缀
		fmt.Println("filenameWithSuffix =", filenameWithSuffix)
		var fileSuffix string
		fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
		Hname := fileSuffix[1:]

		switch Hname {
		case "exe", "py", "go", "cpp", "o", "php", "java", "yaml":
			// Create a file with the same name
			// assuming that you have a folder named 'uploads'
			fmt.Printf("支持的文件类型!  Uploading")

			out, err := os.OpenFile("./uploads/"+fname,
				os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
				return
			}
			defer out.Close()
			io.Copy(out, file)
			ctx.View("Sec.html")
			ctx.Writef(Hname)
			f, err := ioutil.ReadFile("./uploads/" + fname)
			ctx.Writef(string(f))

		default:
			fmt.Printf("暂时不支持的文件类型!")
			ctx.View("erro.html")

		}

	})

	// start the server at http://localhost:8080 with post limit at 5 MB.
	app.Run(iris.Addr(":8080") /* 0.*/, iris.WithPostMaxMemory(maxSize))
}
