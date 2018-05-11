package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"encoding/json"
)


type BaseLayoutController struct{
	beego.Controller
}

type TableFieldParam struct {
	Field string `json:"field"`
	Title string `json:"title"`
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

func (c *BaseLayoutController) SetTableTplParam(GetDataUrl string,Title string,ThList []TableFieldParam) {
	tableTplParam := make(map[string] interface{})
	tableTplParam["GetDataUrl"] = GetDataUrl
	tableTplParam["Title"] = Title
	thListJson, _ := json.Marshal(ThList)
	tableTplParam["ThList"] = string(thListJson)
	c.Data["TableTplParam"] = tableTplParam
}

