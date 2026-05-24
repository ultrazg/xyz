package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ultrazg/xyz/utils"
)

type LoginWithSMSRequestBody struct {
	AreaCode          string `form:"areaCode" json:"areaCode"`
	VerifyCode        string `form:"verifyCode" json:"verifyCode"`
	MobilePhoneNumber string `form:"mobilePhoneNumber" json:"mobilePhoneNumber"`
}

type LoginWithSMSResponseBody struct {
	Data struct {
		User struct {
			Type   string `json:"type"`
			UID    string `json:"uid"`
			Avatar struct {
				Picture struct {
					PicURL       string `json:"picUrl"`
					LargePicURL  string `json:"largePicUrl"`
					MiddlePicURL string `json:"middlePicUrl"`
					SmallPicURL  string `json:"smallPicUrl"`
					ThumbnailURL string `json:"thumbnailUrl"`
					Format       string `json:"format"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
				} `json:"picture"`
			} `json:"avatar"`
			Nickname      string `json:"nickname"`
			IsNicknameSet bool   `json:"isNicknameSet"`
			Bio           string `json:"bio"`
			Gender        string `json:"gender"`
			IsCancelled   bool   `json:"isCancelled"`
			ReadTrackInfo struct {
			} `json:"readTrackInfo"`
			IPLoc          string `json:"ipLoc"`
			BirthYear      int    `json:"birthYear"`
			Industry       string `json:"industry"`
			WechatUserInfo struct {
				NickName string `json:"nickName"`
			} `json:"wechatUserInfo"`
			PhoneNumber struct {
				MobilePhoneNumber string `json:"mobilePhoneNumber"`
				AreaCode          string `json:"areaCode"`
			} `json:"phoneNumber"`
			PhoneNumberNeeded bool `json:"phoneNumberNeeded"`
			JikeUserInfo      struct {
				Nickname string `json:"nickname"`
			} `json:"jikeUserInfo"`
			IsBlockedByViewer bool `json:"isBlockedByViewer"`
		} `json:"user"`
	} `json:"data"`
}

// Login 登录认证
var Login = func(ctx *gin.Context) {
	var params LoginWithSMSRequestBody

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

	p := map[string]any{
		"areaCode":          params.AreaCode,
		"verifyCode":        params.VerifyCode,
		"mobilePhoneNumber": params.MobilePhoneNumber,
	}

	url := "https://podcaster-api.xiaoyuzhoufm.com/v1/auth/login-with-sms"
	headers := map[string]string{
		"accept":          "application/json, text/plain, */*",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"content-type":    "application/json;charset=UTF-8",
		"origin":          "https://podcaster.xiaoyuzhoufm.com",
		"referer":         "https://podcaster.xiaoyuzhoufm.com/",
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36 Edg/146.0.0.0",
	}
	res, code, err := utils.Request(url, http.MethodPost, p, headers)
	if err != nil {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  utils.GetMsg(code),
			"data": err.Error(),
		})

		log.Println(url, code, utils.GetMsg(code))

		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusBadGateway,
			"msg":  utils.GetMsg(http.StatusBadGateway),
			"data": "failed to read upstream login response",
		})

		return
	}

	var data LoginWithSMSResponseBody
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error parsing response body:", err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code": http.StatusBadGateway,
			"msg":  utils.GetMsg(http.StatusBadGateway),
			"data": "failed to parse upstream login response",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  utils.GetMsg(http.StatusOK),
		"data": gin.H{
			"data":                 data.Data.User,
			"x-jike-access-token":  res.Header.Get("x-jike-access-token"),
			"x-jike-refresh-token": res.Header.Get("x-jike-refresh-token"),
		},
	})
}
