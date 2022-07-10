package util

import (
	"github.com/gin-gonic/gin"
	"github.com/hsuehly/presideService/types"
)

// Result 返回的结果：
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(types.ApiCode.SUCCESS)
	res.Msg = types.ApiCode.GetMessage(types.ApiCode.SUCCESS)
	res.Data = data

	c.JSON(200, res)
}

// Error 出错
func Error(c *gin.Context, code int, msg string) {
	res := Result{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}

	c.JSON(200, res)
}
