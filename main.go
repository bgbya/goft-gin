package main

import (
	"github.com/bgbya/goft-gin/src/controllers"
	"github.com/bgbya/goft-gin/src/middlewares"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().
		Attach(middlewares.NewTokenCheck()).
		Mount("v1", controllers.NewIndexController()).
		Launch()
}
