package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wxcloudrun-golang/conf"
)

// 申请二维码
func ApplycodeHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data, _ = getStableAccessToken()
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

type AccessTokenResult struct {
	ErrCode     int    `json:"errcode,omitempty"`
	ErrMsg      string `json:"errmsg,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// getStableAccessToken 获取稳定版接口调用凭据
func getStableAccessToken() (*AccessTokenResult, error) {

	jsonStr := []byte(`{ "grant_type": "client_credential", "appid": "` + conf.AppId + `", "secret": "` + conf.AppSecret + `" }`)
	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/stable_token", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	result := &AccessTokenResult{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
