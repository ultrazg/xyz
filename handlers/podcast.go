package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/utils"
)

type PodcastDetailRequestBody struct {
	Pid string `json:"pid" form:"pid"`
}

// PodcastDetail 查询节目详情
var PodcastDetail = func(ctx *gin.Context) {
	var params PodcastDetailRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Pid == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/podcast/get?pid=" + params.Pid
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

		log.Println("/v1/podcast/get?pid="+params.Pid, code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type RelatedPodcastListRequestBody struct {
	Pid string `json:"pid" form:"pid"`
}

// RelatedPodcastList 相关节目推荐
var RelatedPodcastList = func(ctx *gin.Context) {
	var params RelatedPodcastListRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Pid == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	p := map[string]any{
		"pid":      params.Pid,
		"position": "BOTTOM",
	}
	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/related-podcast/list"
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
		"Content-type":                "application/json",
	}

	response, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/related-podcast/list", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type OwnedPodcastsListBody struct {
	Uid string `form:"uid"`
}

// OwnedPodcastsList 用户创建的播客
var OwnedPodcastsList = func(ctx *gin.Context) {
	var params *OwnedPodcastsListBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	if params.Uid == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")
	p := map[string]any{
		"uid": params.Uid,
	}
	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/podcaster/owned-podcasts"
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
		"Content-type":                "application/json",
	}

	response, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/podcaster/owned-podcasts", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type PodcastGetInfoBody struct {
	Pid string `form:"pid"`
}

// PodcastGetInfo 获取节目主体信息
var PodcastGetInfo = func(ctx *gin.Context) {
	var params *PodcastGetInfoBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Pid == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	url := constant.BaseUrl + "/v1/podcast/get-info?pid=" + params.Pid
	headers := map[string]string{
		"Host":                "api.xiaoyuzhoufm.com",
		"User-Agent":          "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1) xyzTheme/defaultLight",
		"Referer":             "https://terms.xiaoyuzhoufm.com/",
		"Market":              "AppStore",
		"App-BuildNo":         "1576",
		"OS":                  "ios",
		"Origin":              "https://terms.xiaoyuzhoufm.com",
		"Sec-Fetch-Dest":      "empty",
		"Sec-Fetch-Site":      "same-site",
		"Sec-Fetch-Mode":      "cors",
		"x-jike-access-token": XJikeAccessToken,
		"Manufacturer":        "Apple",
		"BundleID":            "app.podcast.cosmos",
		"Connection":          "keep-alive",
		"Accept-Language":     "zh-CN,zh-Hans;q=0.9",
		"Model":               "iPhone14,2",
		"Accept":              "application/json, text/plain, */*",
		"App-Version":         "2.57.1",
		"OS-Version":          "17.4.1",
	}

	response, code, err := utils.Request(url, http.MethodGet, nil, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/podcast/get-info?pid="+params.Pid, code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type PodcastHonorListBody struct {
	Pid string `form:"pid"`
}

// PodcastHonorList 获取节目荣誉墙
var PodcastHonorList = func(ctx *gin.Context) {
	var params *PodcastHonorListBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	if params.Pid == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")
	p := map[string]any{
		"pid": params.Pid,
	}
	url := constant.BaseUrl + "/v1/podcast-honor/list"
	headers := map[string]string{
		"Host":                "api.xiaoyuzhoufm.com",
		"User-Agent":          "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Xiaoyuzhou/2.57.1 (build:1576; iOS 17.4.1) xyzTheme/defaultLight",
		"Referer":             "https://h5.xiaoyuzhoufm.com/",
		"Origin":              "https://h5.xiaoyuzhoufm.com/",
		"Market":              "AppStore",
		"App-BuildNo":         "1576",
		"Sec-Fetch-Dest":      "empty",
		"Sec-Fetch-Mode":      "cors",
		"OS":                  "ios",
		"x-jike-access-token": XJikeAccessToken,
		"Sec-Fetch-Site":      "same-site",
		"Manufacturer":        "Apple",
		"BundleID":            "app.podcast.cosmos",
		"Connection":          "keep-alive",
		"Accept-Language":     "zh-CN,zh-Hans;q=0.9",
		"Model":               "iPhone14,2",
		"Accept":              "application/json, text/plain, */*",
		"App-Version":         "2.57.1",
		"OS-Version":          "17.4.1",
		"Content-type":        "application/json",
	}

	response, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/podcast-honor/list", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
