package controllers

type MainController struct {
	BaseLayoutController
}

func (c *MainController) Get() {
	c.Data["BodyClass"] = "hold-transition sidebar-mini"
	thList := []TableFieldParam{
		{Title: "姓名", Field: "PlayerName"},
		{Title: "手机号码", Field: "Demo"},
		{Title: "爱好", Field: "DemoOk"},
		{Title: "地址", Field: "Demottt"},
	}
	c.SetTableTplParam("/v1/object","DemoTitle",thList)
	c.TplName = "index.html"
}