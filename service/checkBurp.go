package service

import (
	"checkBurp/model"
	"github.com/kataras/iris"
	"strings"
)

func RemoteInfo(ctx iris.Context)(model.RemoteAdd)  {
	remote :=model.RemoteAdd{}
	ip:=ctx.Request().RemoteAddr
	if len(strings.Split(ip,":"))==2{
		remote.RemoteIp=strings.Split(ip,":")[0]
		remote.RemotePort=strings.Split(ip,":")[1]
		remote.Url=ctx.Scheme()+ctx.Host()+"/second"
		return remote
	}else {
		remote.RemoteIp="null"
		remote.RemoteIp="null"
		remote.Url="null"
		return remote
	}

}

func SecondInfo(ctx iris.Context)(model.RemoteAdd)  {
	remote :=model.RemoteAdd{}
	ip:=ctx.Request().RemoteAddr
	if len(strings.Split(ip,":"))==2{
		remote.RemoteIp=strings.Split(ip,":")[0]
		remote.RemotePort=strings.Split(ip,":")[1]
		return remote
	}else {
		remote.RemoteIp="null"
		remote.RemoteIp="null"
		return remote
	}

}