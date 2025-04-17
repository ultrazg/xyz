package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/utils"
)

type InboxListRequestBody struct {
	LoadMoreKey *InboxListLoadMoreKey `json:"loadMoreKey" form:"loadMoreKey"`
}

type InboxListLoadMoreKey struct {
	PubDate string `json:"pubDate" form:"pubDate"`
	Id      string `json:"id" form:"id"`
}

// InboxList 订阅更新列表
var InboxList = func(ctx *gin.Context) {

	var params *InboxListRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	p := map[string]any{
		"limit": "20",
	}

	if params.LoadMoreKey != nil {
		p["loadMoreKey"] = map[string]string{
			"pubDate": params.LoadMoreKey.PubDate,
			"id":      params.LoadMoreKey.Id,
		}
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/inbox/list"
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
		"Accept-Language":             "zh-Hans-CN;q=1.0, zh-Hant-TW;q=0.9",
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
		"x-jike-device-id":            "",
		"x-jike-device-properties":    "",
	}

	response, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/inbox/list", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
