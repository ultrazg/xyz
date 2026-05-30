package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/constant"
	"github.com/ultrazg/xyz/utils"
)

type privateMediaRequestBody struct {
	Eid string `form:"eid"`
}

var GetPrivateMedia = func(ctx *gin.Context) {
	var params privateMediaRequestBody

	err := ctx.ShouldBind(&params)
	if err != nil {
		utils.ReturnBadRequest(ctx, err)
		return
	}

	h := ctx.Request.Header
	XJikeAccessToken := h.Get("x-jike-access-token")

	if params.Eid == "" {
		utils.ReturnBadRequest(ctx, nil)
		return
	}

	now := time.Now()
	isoTime := now.Format("2006-01-02T15:04:05Z07:00")
	url := constant.BaseUrl + "/v1/private-media/get?dubbing=false&eid=" + params.Eid
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
		"Model":                       "iPhone14,2",
		"app-permissions":             "100000",
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
		log.Println("/v1/private-media/get", code, utils.GetMsg(code))
		return
	}

	utils.ReturnJson(response, ctx)
}
