package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/utils"
	"log"
	"net/http"
	"time"
)

// GetMileage 获取收听数据概览
var GetMileage = func(ctx *gin.Context) {
	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")
	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/mileage/get"
	headers := map[string]string{
		"Host":                        "api.xiaoyuzhoufm.com",
		"User-Agent":                  "Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1)",
		"Market":                      "AppStore",
		"App-BuildNo":                 "1576",
		"OS":                          "ios",
		"x-jike-access-token":         XJikeAccessToken,
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
		"WifiConnected":               "true",
		"OS-Version":                  "17.4.1",
		"x-custom-xiaoyuzhou-app-dev": "",
		"Local-Time":                  isoTime,
		"Timezone":                    "Asia/Shanghai",
	}

	response, code, err := utils.Request(url, http.MethodGet, nil, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/mileage/get", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type MileageBody struct {
	All bool `form:"all"`
}

// GetMileageList 获取收听数据排行榜
var GetMileageList = func(ctx *gin.Context) {
	var params *MileageBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	p := map[string]any{}

	if params.All {
		p["rank"] = "TOTAL"
	} else {
		p["rank"] = "LAST_THIRTY_DAYS"
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")
	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/mileage/list"
	headers := map[string]string{
		"Host":                        "api.xiaoyuzhoufm.com",
		"User-Agent":                  "Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1)",
		"Market":                      "AppStore",
		"App-BuildNo":                 "1576",
		"OS":                          "ios",
		"x-jike-access-token":         XJikeAccessToken,
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
		"WifiConnected":               "true",
		"OS-Version":                  "17.4.1",
		"x-custom-xiaoyuzhou-app-dev": "",
		"Local-Time":                  isoTime,
		"Timezone":                    "Asia/Shanghai",
	}

	response, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/mileage/list", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
