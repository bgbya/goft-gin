package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type IndexController struct{}

func (ic *IndexController) Name() string {
	return "ic"
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (ic *IndexController) IndexTest(ctx *gin.Context) goft.Json {
	goft.Throw("测试异常", 0, ctx)
	return gin.H{"result": "index test"}
}

func (ic *IndexController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", ic.IndexTest)
}
