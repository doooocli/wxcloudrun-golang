package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wxcloudrun-golang/conf"
)

// ApplycodeHandler 申请二维码
func ApplycodeHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}

	result, err := marketcodeApplycode()
	res.Data = result
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintln(w, err)
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

	//resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + conf.AppId + "&secret=" + conf.AppSecret)

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

type Applycode struct {
	ErrCode       int    `json:"errcode,omitempty"`
	ErrMsg        string `json:"errmsg,omitempty"`
	ApplicationId string `json:"application_id,omitempty"`
}

// marketcodeApplycode 申请二维码接口
func marketcodeApplycode() (*Applycode, error) {

	jsonStr := []byte(`{"code_count": 10000, "isv_application_id": "order001"}`)
	resp, err := http.Post("http://api.weixin.qq.com/intp/marketcode/applycode", "application/json", bytes.NewBuffer(jsonStr))

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
	result := &Applycode{}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
