 package main

 import (
	 "github.com/kataras/iris"
	 "github.com/kataras/iris/middleware/logger"
	 "github.com/kataras/iris/middleware/recover"
	 "checkBurp/control"
 )

 func router(app  *iris.Application)  {

      //app.Get("/", func(ctx iris.Context) {
      //    ctx.HTML(" <h1>hi, I just exist in order to see if the server is closed</h1>")
      //})
	 app.Get("/", func(ctx iris.Context) {
	 	ctx.HTML("<h1>welcole</h1>")

	 })
      apis:=app.Party("/")
      {
      	apis.Get("/remoteInfo",control.RemoteInfo)
      	apis.Get("/second",control.SecondInfo)

	 }

 }
    func main() {
        app := iris.New()
        app.Use(recover.New())
        app.Use(logger.New())
        app.Logger().SetLevel("info")
        router(app)
        app.Run(iris.Addr("0.0.0.0:8089"), iris.WithoutInterruptHandler)

    }