package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/pkg/utils"
	"log"
	"net/http"
	"time"
)

type EpisodePlayedHistoryListRequestBody struct {
	LoadMoreKey string `form:"loadMoreKey"`
}

// EpisodePlayedHistoryList 收听历史
var EpisodePlayedHistoryList = func(ctx *gin.Context) {
	var params *EpisodePlayedHistoryListRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	p := map[string]any{}

	if params.LoadMoreKey != "" {
		p["loadMoreKey"] = params.LoadMoreKey
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/episode-played/list-history"
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

		log.Println("/v1/episode-played/list-history", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
