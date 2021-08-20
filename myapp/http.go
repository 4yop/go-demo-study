package main

import "github.com/kataras/iris/v12"

func main()  {
	app := iris.Default()

	// This handler will match /user/john but will not match /user/ or /user
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})
	app.Listen(":8080")
}