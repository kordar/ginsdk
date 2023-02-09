package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// 成功
	success = 200
	// 数据列表
	datalist = 200
	excel    = 200
	fail     = 500
	// 异常告警
	warn   = 501
	refuse = 401
	// 登录失败
	authFail = 402
	// 数据校验异常
	validate = 3000
)

func Result(c *gin.Context, code int, message string, data interface{}, count int64) {
	if data == nil {
		c.JSON(http.StatusOK, map[string]interface{}{"code": code, "message": message})
	} else {
		if count == -1 {
			c.JSON(http.StatusOK, map[string]interface{}{"code": code, "message": message, "data": data})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{"code": code, "message": message, "data": data, "count": count})
		}
	}
}

func Success(c *gin.Context, message string, data interface{}) {
	Result(c, success, message, data, -1)
}

func Fail(c *gin.Context, message string, data interface{}) {
	Result(c, fail, message, data, -1)
}

func Warn(c *gin.Context, message string, data interface{}) {
	Result(c, warn, message, data, -1)
}

func Data(c *gin.Context, message string, data interface{}, count int64) {
	Result(c, datalist, message, data, count)
}

func SuccessOrWarn(c *gin.Context, flag bool, successMessage string, failMessage string) {
	if flag {
		Result(c, success, successMessage, nil, -1)
	} else {
		Result(c, warn, failMessage, nil, -1)
	}
}

func Excel(c *gin.Context, data interface{}, header interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": success, "message": "success", "data": data, "header": header,
	})
}
