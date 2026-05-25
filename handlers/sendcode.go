package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/utils"
)

type SendCodeRequestBody struct {
	MobilePhoneNumber string `form:"mobilePhoneNumber" json:"mobilePhoneNumber"`
	AreaCode          string `form:"areaCode" json:"areaCode"`
}

// SendCode 发送短信验证码
var SendCode = func(ctx *gin.Context) {
	var params SendCodeRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	if params.MobilePhoneNumber == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	if params.AreaCode == "" {
		params.AreaCode = "+86"
	}

	p := map[string]any{
		"mobilePhoneNumber": params.MobilePhoneNumber,
		"areaCode":          params.AreaCode,
	}
	url := "https://podcaster-api.xiaoyuzhoufm.com/v1/auth/send-code"
	headers := map[string]string{
		"accept":          "application/json, text/plain, */*",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"content-type":    "application/json;charset=UTF-8",
		"origin":          "https://podcaster.xiaoyuzhoufm.com",
		"referer":         "https://podcaster.xiaoyuzhoufm.com/",
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36 Edg/146.0.0.0",
	}

	res, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println(url, code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(res, ctx)
}
