package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/pkg/utils"
	"io"
	"log"
	"net/http"
)

type LoginOrSignUpWithSMSRequestBody struct {
	AreaCode          string `form:"areaCode" json:"areaCode"`
	VerifyCode        string `form:"verifyCode" json:"verifyCode"`
	MobilePhoneNumber string `form:"mobilePhoneNumber" json:"mobilePhoneNumber"`
}

// Login 登录认证
var Login = func(ctx *gin.Context) {
	var params LoginOrSignUpWithSMSRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	if params.VerifyCode == "" || params.MobilePhoneNumber == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	if params.AreaCode == "" {
		params.AreaCode = "+86"
	}

	p := map[string]string{
		"areaCode":          params.AreaCode,
		"verifyCode":        params.VerifyCode,
		"mobilePhoneNumber": params.MobilePhoneNumber,
	}

	url := constant.BaseUrl + "/v1/auth/loginOrSignUpWithSMS"
	headers := map[string]string{
		"Host":            "api.xiaoyuzhoufm.com",
		"App-BuildNo":     "1576",
		"OS":              "ios",
		"Manufacturer":    "Apple",
		"BundleID":        "app.podcast.cosmos",
		"abtest-info":     "{\"old_user_discovery_feed\":\"enable\"}",
		"Accept-Language": "zh-Hant-HK;q=1.0, zh-Hans-CN;q=0.9",
		"Model":           "iPhone14,2",
		"app-permissions": "4",
		"Accept":          "*/*",
		"Content-Type":    "application/json",
		"App-Version":     "2.57.1",
		"WifiConnected":   "true",
		"OS-Version":      "17.4.1",
	}
	res, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/auth/loginOrSignUpWithSMS", code, utils.GetMsg(code))

		return
	}

	body, err := io.ReadAll(res.Body)
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
		"msg":  utils.GetMsg(http.StatusOK),
		"data": gin.H{
			"data":                 data,
			"x-jike-access-token":  res.Header.Get("x-jike-access-token"),
			"x-jike-refresh-token": res.Header.Get("x-jike-refresh-token"),
		},
	})
}
