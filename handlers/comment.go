package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/pkg/utils"
	"log"
	"net/http"
	"time"
)

type CommentPrimaryRequestBody struct {
	Id          string                     `json:"id" form:"id"`
	Order       string                     `json:"order" form:"order"`
	LoadMoreKey *commentPrimaryLoadMoreKey `json:"loadMoreKey" form:"loadMoreKey"`
}

type commentPrimaryLoadMoreKey struct {
	Id           string  `json:"id" form:"id"`
	Direction    string  `json:"direction" form:"direction"`
	HotSortScore float64 `json:"hotSortScore" form:"hotSortScore"`
}

// CommentPrimary 评论
var CommentPrimary = func(ctx *gin.Context) {
	var params CommentPrimaryRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Order == "" || params.Id == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	p := map[string]any{
		"order": params.Order,
		"owner": map[string]any{
			"id":   params.Id,
			"type": "EPISODE",
		},
	}

	if params.LoadMoreKey != nil {
		p["loadMoreKey"] = map[string]any{
			"hotSortScore": params.LoadMoreKey.HotSortScore,
			"id":           params.LoadMoreKey.Id,
			"direction":    params.LoadMoreKey.Direction,
		}
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/comment/list-primary"
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

		log.Println("/v1/comment/list-primary", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}

type CommentThreadRequestBody struct {
	Order            string `json:"order" form:"order"`
	PrimaryCommentId string `json:"primaryCommentId" form:"primaryCommentId"`
}

// CommentThread 评论回复
var CommentThread = func(ctx *gin.Context) {
	var params CommentThreadRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Order == "" || params.PrimaryCommentId == "" {
		utils.ReturnBadRequest(ctx, nil)

		return
	}

	p := map[string]any{
		"order":            params.Order,
		"primaryCommentId": params.PrimaryCommentId,
	}
	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/comment/list-thread"
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
		"X-Online-Host":               "api.xiaoyuzhoufm.com",
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

		log.Println("/v1/comment/list-thread", code, utils.GetMsg(code))

		return
	}

	utils.ReturnJson(response, ctx)
}
