package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

var MsgFlag = map[int]string{
	200: "OK",
	400: "错误请求",
	401: "身份认证信息失效，尝试重新登录或调用 /refresh_token 接口以获取有效认证信息",
	403: "拒绝访问",
	404: "请求的资源不存在",
	405: "请求方法不支持",
	500: "服务器内部错误",
	502: "网关错误",
	503: "服务不可用",
	504: "网关超时",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}

	return "Error"
}

// ReturnBadRequest 错误参数
func ReturnBadRequest(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
			"msg":   GetMsg(http.StatusBadRequest),
		})

		log.Println(err.Error())
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  GetMsg(http.StatusBadRequest),
		})

		log.Println(GetMsg(http.StatusBadRequest))
	}
}

func ReturnJson(response *http.Response, ctx *gin.Context) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)

		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error parsing response body:", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  GetMsg(http.StatusOK),
		"data": data,
	})
}
