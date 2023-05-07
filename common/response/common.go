package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespEntity struct {
	//0是正確,其他需要參考接口錯誤碼對應
	Code int `json:"code"`
	//接口数据
	Data interface{} `json:"data"`
	//响应消息
	Msg string `json:"msg"`
}

func Result(code ErrorCode, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, RespEntity{
		int(code),
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(0, map[string]interface{}{}, "操作成功", c)
}
func Error(err error, c *gin.Context) {
	e, ok := err.(AgamottoError)
	if ok {
		c.JSON(http.StatusOK, RespEntity{
			int(e.GetCode()),
			"",
			err.Error(),
		})
		return
	}
	Result(10001, gin.H{}, err.Error(), c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(0, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(0, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(0, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(SystemError, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SystemError, data, message, c)
}

func ResultCodeMessage(code ErrorCode, message string, c *gin.Context) {
	if code == 0 {
		OkWithData(map[string]interface{}{}, c)
	} else {
		Result(code, map[string]interface{}{}, message, c)
	}
}
