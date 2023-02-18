package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TokenCheck struct {
}

func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (t TokenCheck) OnRequest(context *gin.Context) error {
	if context.Query("token") == "" {
		return fmt.Errorf("token不能为空")
	}
	return nil
}

func (t TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	panic("implement me")
}
