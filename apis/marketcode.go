package apis

import (
	"encoding/json"
	"net/http"
	"wxcloudrun-golang/service/marketcode"
	"wxcloudrun-golang/util/response"
)

// ApplyCodeHandler 申请二维码
func ApplyCodeHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	result, err := marketcode.GetApplyCode(param)
	if err != nil {
		return
	}
	response.NewResponse(w).OK(result)
}

// ApplyCodeQueryHandler 申请二维码查询
func ApplyCodeQueryHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	result, err := marketcode.ApplyCodeQuery(param)
	if err != nil {
		return
	}
	response.NewResponse(w).OK(result)
}

// GetApplyCodeDownloadHandler 二维码下载
func GetApplyCodeDownloadHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	result, err := marketcode.GetApplyCodeDownload(param)
	if err != nil {
		return
	}
	response.NewResponse(w).OK(result)
}

// CodeActiveHandler 二维码激活
func CodeActiveHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.CodeActiveParam{}
	json.NewDecoder(r.Body).Decode(param)
	result, err := marketcode.CodeActive(param)
	if err != nil {
		return
	}
	response.NewResponse(w).OK(result)
}

// CodeActiveQueryHandler 二维码激活查询
func CodeActiveQueryHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	result, err := marketcode.CodeActiveQuery(param)
	if err != nil {
		return
	}
	response.NewResponse(w).OK(result)
}
