package handlers

import (
	"github.com/gin-gonic/gin"
	C "github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/utils"
	"log"
	"net/http"
)

type RefreshTokenRequestBody struct {
	XJikeAccessToken  string `json:"x-jike-access-token" form:"x-jike-access-token"`
	XJikeRefreshToken string `json:"x-jike-refresh-token" form:"x-jike-refresh-token"`
}

// RefreshToken 刷新token
var RefreshToken = func(ctx *gin.Context) {
	var params RefreshTokenRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	if params.XJikeAccessToken == "" || params.XJikeRefreshToken == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	url := C.BaseUrl + "/app_auth_tokens.refresh"
	headers := map[string]string{
		"Host":                        "api.xiaoyuzhoufm.com",
		"User-Agent":                  "Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1)",
		"Market":                      "AppStore",
		"App-BuildNo":                 "1576",
		"OS":                          "ios",
		"x-jike-access-token":         params.XJikeAccessToken,
		"x-jike-refresh-token":        params.XJikeRefreshToken,
		"Manufacturer":                "Apple",
		"BundleID":                    "app.podcast.cosmos",
		"Connection":                  "keep-alive",
		"Accept-Language":             "zh-Hant-HK;q=1.0, zh-Hans-CN;q=0.9",
		"Model":                       "iPhone14,2",
		"app-permissions":             "4",
		"Accept":                      "*/*",
		"Content-Type":                "application/x-www-form-urlencoded; charset=utf-8",
		"App-Version":                 "2.57.1",
		"WifiConnected":               "true",
		"OS-Version":                  "17.4.1",
		"x-custom-xiaoyuzhou-app-dev": "",
	}

	response, code, err := utils.Request(url, http.MethodPost, nil, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/app_auth_tokens.refresh", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
