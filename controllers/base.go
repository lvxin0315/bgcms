package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)


type BaseLayoutController struct{
	beego.Controller
}

func (c *BaseLayoutController) Init (ctx *context.Context, controllerName, actionName string, app interface{})  {
	c.Layout = "layout/layout.html"
	c.TplName = ""
	c.Ctx = ctx
	c.TplExt = "tpl"
	c.AppController = app
	c.EnableRender = true
	c.EnableXSRF = true
	c.Data = ctx.Input.Data()

}

