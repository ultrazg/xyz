package handlers

import (
	"github.com/gin-gonic/gin"
	C "github.com/ultrazg/xyz/constant"
	U "github.com/ultrazg/xyz/pkg/utils"
	"log"
	"net/http"
	"time"
)

type SearchRequestBody struct {
	Type        string             `json:"type" form:"type"`
	Keyword     string             `json:"keyword" form:"keyword"`
	LoadMoreKey *searchLoadMoreKey `json:"loadMoreKey" form:"loadMoreKey"`
}

type searchLoadMoreKey struct {
	LoadMoreKey int    `json:"loadMoreKey" form:"loadMoreKey"`
	SearchId    string `json:"searchId" form:"searchId"`
}

// Search 搜索
var Search = func(ctx *gin.Context) {
	var params SearchRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		U.ReturnBadRequest(ctx, err)

		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if XJikeAccessToken == "" || params.Keyword == "" || params.Type == "" {
		U.ReturnBadRequest(ctx, nil)

		return
	}

	p := map[string]any{
		"limit":           "20",
		"sourcePageName":  "4",
		"type":            params.Type,
		"currentPageName": "4",
		"keyword":         params.Keyword,
	}

	if params.LoadMoreKey != nil {
		p["loadMoreKey"] = map[string]any{
			"loadMoreKey": params.LoadMoreKey.LoadMoreKey,
			"searchId":    params.LoadMoreKey.SearchId,
		}
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := C.BaseUrl + "/v1/search/create"
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

	res, code, err := U.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  U.GetMsg(code),
			"data": err.Error(),
		})

		log.Println("/v1/search/create", code, U.GetMsg(code))

		return
	}

	U.ReturnJson(res, ctx)
}
