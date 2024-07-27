package handlers

import (
	"github.com/ultrazg/xyz/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	C "github.com/ultrazg/xyz/constant"
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

	url := C.BaseUrl + "/v1/auth/sendCode"
	headers := map[string]string{
		"Host":                        "api.xiaoyuzhoufm.com",
		"User-Agent":                  "Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1)",
		"Market":                      "AppStore",
		"App-BuildNo":                 "1576",
		"OS":                          "ios",
		"Manufacturer":                "Apple",
		"BundleID":                    "app.podcast.cosmos",
		"Connection":                  "keep-alive",
		"abtest-info":                 "{\"old_user_discovery_feed\":\"enable\"}",
		"Accept-Language":             "zh-Hant-HK;q=1.0, zh-Hans-CN;q=0.9",
		"Model":                       "iPhone14,2",
		"app-permissions":             "4",
		"Accept":                      "*/*",
		"Content-Type":                "application/json",
		"App-Version":                 "2.57.1",
		"Accept-Encoding":             "br;q=1.0, gzip;q=0.9, deflate;q=0.8",
		"WifiConnected":               "true",
		"OS-Version":                  "17.4.1",
		"x-custom-xiaoyuzhou-app-dev": "",
	}

	res, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/auth/sendCode", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(res, ctx)
}
