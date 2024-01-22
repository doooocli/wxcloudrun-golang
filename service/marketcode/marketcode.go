package marketcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wxcloudrun-golang/conf"
)

type ApplyCodeParam struct {
	CodeCount        int    `json:"code_count,omitempty"`
	IsvApplicationId string `json:"isv_application_id,omitempty"`
	ApplicationId    string `json:"application_id,omitempty"`
	CodeStart        int    `json:"code_start,omitempty"`
	CodeEnd          int    `json:"code_end,omitempty"`
}

type CodeActiveParam struct {
	ApplyCodeParam
	ActivityName string `json:"activity_name,omitempty"` //活动名称
	ProductBrand string `json:"product_brand,omitempty"` //商品品牌
	ProductTitle string `json:"product_title,omitempty"` //商品标题
	ProductCode  string `json:"product_code,omitempty"`  //商品条码
	WxaAppid     string `json:"wxa_appid,omitempty"`     //小程序的appid
	WxaPath      string `json:"wxa_path,omitempty"`      //小程序的path
	WxaType      int    `json:"wxa_type,omitempty"`      //小程序版本
}

type TicketToCodeParam struct {
	OpenId     int    `json:"openid"`
	CodeTicket string `json:"code_ticket"`
}

type Resp struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type AccessTokenResp struct {
	Resp
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

// GetStableAccessToken 获取稳定版接口调用凭据
func GetStableAccessToken() (*AccessTokenResp, error) {
	param := &struct {
		grant_type string
		appid      string
		secret     string
	}{
		grant_type: "client_credential",
		appid:      conf.AppId,
		secret:     conf.AppSecret,
	}
	result := &AccessTokenResp{}
	if err := httpSend("https://api.weixin.qq.com/cgi-bin/stable_token", param, result); err != nil {
		return nil, err

	}
	return result, nil
}

type ApplyCodeResp struct {
	Resp
	ApplicationId string `json:"application_id,omitempty"`
}

// GetApplyCode 申请二维码接口
func GetApplyCode(param *ApplyCodeParam) (*ApplyCodeResp, error) {
	result := &ApplyCodeResp{}
	if err := httpSend("http://api.weixin.qq.com/intp/marketcode/applycode", param, result); err != nil {
		return nil, err
	}
	return result, nil
}

type ApplyCodeQueryResp struct {
	Resp
	Status           string `json:"status,omitempty"`
	IsvApplicationId string `json:"isv_application_id,omitempty"`
	ApplicationId    int    `json:"application_id,omitempty"`
	CreateTime       int    `json:"create_time,omitempty"`
	UpdateTime       int    `json:"update_time,omitempty"`
	CodeStart        int    `json:"code_start,omitempty"`
	CodeEnd          int    `json:"code_end,omitempty"`
}

// ApplyCodeQuery 查询二维码申请单接口
func ApplyCodeQuery(param *ApplyCodeParam) (*ApplyCodeQueryResp, error) {
	result := &ApplyCodeQueryResp{}
	if err := httpSend("http://api.weixin.qq.com/intp/marketcode/applycodequery", param, result); err != nil {
		return nil, err
	}
	return result, nil
}

type ApplyCodeDownloadResp struct {
	Resp
	Buffer string `json:"buffer,omitempty"`
}

// GetApplyCodeDownload 下载二维码包接口
func GetApplyCodeDownload(param *ApplyCodeParam) (*ApplyCodeDownloadResp, error) {
	result := &ApplyCodeDownloadResp{}
	if err := httpSend("http://api.weixin.qq.com/intp/marketcode/applycodedownload", param, result); err != nil {
		return nil, err
	}

	// TODO 解密
	return result, nil
}

// CodeActive 激活二维码接口
func CodeActive(param *CodeActiveParam) (*Resp, error) {
	result := &Resp{}
	if err := httpSend("http://api.weixin.qq.com/intp/marketcode/codeactive", param, result); err != nil {
		return nil, err
	}
	return result, nil
}

type CodeActiveQueryResp struct {
	Resp
	Code             string `json:"code,omitempty"`
	ApplicationId    string `json:"application_id,omitempty"`
	IsvApplicationId string `json:"isv_application_id,omitempty"`
	ActivityName     string `json:"activity_name,omitempty"`
	ProductBrand     string `json:"product_brand,omitempty"`
	ProductTitle     string `json:"product_title,omitempty"`
	WxaAppid         string `json:"wxa_appid,omitempty"`
	WxaPath          string `json:"wxa_path,omitempty"`
	WxaType          string `json:"wxa_type,omitempty"`
	CodeStart        string `json:"code_start,omitempty"`
	CodeEnd          string `json:"code_end,omitempty"`
}

// CodeActiveQuery 查询二维码激活状态接口
func CodeActiveQuery(param *ApplyCodeParam) (*CodeActiveQueryResp, error) {
	result := &CodeActiveQueryResp{}
	if err := httpSend("http://api.weixin.qq.com/intp/marketcode/codeactivequery", param, result); err != nil {
		return nil, err
	}
	return result, nil
}

// TicketToCode code_ticket换code
func TicketToCode(param *TicketToCodeParam) (*CodeActiveQueryResp, error) {
	result := &CodeActiveQueryResp{}
	if err := httpSend("https://api.weixin.qq.com/intp/marketcode/tickettocode", param, result); err != nil {
		return nil, err
	}
	return result, nil
}

func httpSend(url string, param interface{}, to interface{}) error {
	paramStr, err := json.Marshal(param)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(paramStr))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, to); err != nil {
		return err
	}
	return nil
}
