package control

import (
	"checkBurp/service"
	"checkBurp/utils"
	"github.com/kataras/iris"
)

func RemoteInfo(ctx iris.Context)  {
	a := service.RemoteInfo(ctx)
	ctx.JSON(utils.ResultUtil.Success(a))

}

func SecondInfo(ctx iris.Context)  {
	a := service.SecondInfo(ctx)
	ctx.JSON(utils.ResultUtil.Success(a))

}
