package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Dct map[string]interface {}

func (c App) Index() revel.Result {
	greeting := "Test String.."
	return c.Render(greeting)
}


func (c App) HelloWorld() revel.Result{
	tt:=Dct{
		"name":"Zhangsan!",
		"address":"City of Peking",
	}
	return c.RenderJSON(tt)
}

func (c App) AutoLogin() revel.Result{
	return c.Render()
}