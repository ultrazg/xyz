package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ultrazg/xyz/constant"
	"io"
	"net/http"
	"strings"
)

type ResponseBody struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

func CheckUpgrade() error {
	response, _, err := Request(constant.UpgradeUrl, http.MethodGet, nil, nil)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var data *ResponseBody
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	remoteVersion := strings.TrimPrefix(data.TagName, "v")

	if constant.Version != remoteVersion {
		fmt.Printf("\r\n✨发现新版本\r\n当前版本：%s\r\n最新版本：%s\r\n更新内容：\n%s\r\n%s", constant.Version, remoteVersion, data.Body, constant.ReleaseUrl)
	}

	return nil
}
