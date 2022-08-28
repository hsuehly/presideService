package Response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Codes struct {
	SUCCESS           int
	FAILED            int
	GENERATETOKEN     int
	NOAUTH            int
	AUTHFAILED        int
	AUTHFORMATERROR   int
	INVALIDTOKEN      int
	NOSUCHID          int
	CREATEORDERFAILED int
	LCAKPARAMETERS    int
	CONVERTFAILED     int
	NOSUCHNAME        int
	EXISTSNAME        int
	Message           map[int]string
}

var ApiCode = &Codes{
	SUCCESS:           200,
	FAILED:            0,
	AUTHFAILED:        401,
	GENERATETOKEN:     402,
	NOAUTH:            403,
	AUTHFORMATERROR:   404,
	INVALIDTOKEN:      405,
	NOSUCHID:          101,
	CREATEORDERFAILED: 201,
	LCAKPARAMETERS:    301,
	CONVERTFAILED:     302,
	NOSUCHNAME:        501,
	EXISTSNAME:        502,
}

func init() {
	ApiCode.Message = map[int]string{
		ApiCode.SUCCESS:           "success",
		ApiCode.FAILED:            "error",
		ApiCode.GENERATETOKEN:     "生成Token失败",
		ApiCode.AUTHFAILED:        "鉴权失败",
		ApiCode.NOAUTH:            "请求头中auth为空",
		ApiCode.AUTHFORMATERROR:   "请求头中auth格式有误",
		ApiCode.INVALIDTOKEN:      "无效的Token",
		ApiCode.NOSUCHID:          "id不存在",
		ApiCode.CREATEORDERFAILED: "预约订单失败",
		ApiCode.LCAKPARAMETERS:    "缺少参数",
		ApiCode.CONVERTFAILED:     "参数类型转换报错",
		ApiCode.NOSUCHNAME:        "根据名称查不到数据",
		ApiCode.EXISTSNAME:        "名称重复",
	}
}

func (c *Codes) GetMessage(code int) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}

// Result 返回的结果：
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ValidatorResult struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = ApiCode.SUCCESS
	res.Msg = ApiCode.GetMessage(ApiCode.SUCCESS)
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

func ValidatorErr(c *gin.Context, err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		Error(c, ApiCode.CONVERTFAILED, ApiCode.GetMessage(ApiCode.CONVERTFAILED))
		return
	}
	var res ValidatorResult
	res.Code = 400
	res.Msg = removeTopStruct(errs.Translate(trans))
	c.JSON(200, res)
	return
}
