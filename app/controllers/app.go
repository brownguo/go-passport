package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Test String.."
	return c.Render(greeting)
}


func (c App) HelloWorld() revel.Result{
	return c.Render()
}
