package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/pkg/utils"
	"log"
	"net/http"
	"time"
)

type EpisodeListRequestBody struct {
	Pid string `json:"pid" form:"pid"`
}

// EpisodeList 查询节目列表
var EpisodeList = func(ctx *gin.Context) {
	var params EpisodeListRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Pid == "" || XJikeAccessToken == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/episode/list"
	p := map[string]string{
		// TODO: 接收 loadMoreKey 对象用于翻页
		"limit": "20",
		"pid":   params.Pid,
	}
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

		log.Println("/v1/episode/list", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type EpisodeDetailRequestBody struct {
	Eid string `json:"eid" form:"eid"`
}

// EpisodeDetail 查询单集的详情
var EpisodeDetail = func(ctx *gin.Context) {
	var params EpisodeDetailRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Eid == "" || XJikeAccessToken == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/episode/get?eid=" + params.Eid
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

		log.Println("/v1/episode/get?eid="+params.Eid, code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
