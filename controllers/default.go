package controllers

type MainController struct {
	BaseLayoutController
}

func (c *MainController) Get() {
	c.Data["BodyClass"] = "hold-transition sidebar-mini"
	c.TplName = "index.html"
}
