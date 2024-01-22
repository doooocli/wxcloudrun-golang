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
	if result, err := marketcode.GetApplyCode(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}

// ApplyCodeQueryHandler 申请二维码查询
func ApplyCodeQueryHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	if result, err := marketcode.ApplyCodeQuery(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}

// GetApplyCodeDownloadHandler 二维码下载
func GetApplyCodeDownloadHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	if result, err := marketcode.GetApplyCodeDownload(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}

// CodeActiveHandler 二维码激活
func CodeActiveHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.CodeActiveParam{}
	json.NewDecoder(r.Body).Decode(param)
	if result, err := marketcode.CodeActive(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}

// CodeActiveQueryHandler 二维码激活查询
func CodeActiveQueryHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.ApplyCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	if result, err := marketcode.CodeActiveQuery(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}

// TicketToCodeHandler 二维码激活查询
func TicketToCodeHandler(w http.ResponseWriter, r *http.Request) {
	param := &marketcode.TicketToCodeParam{}
	json.NewDecoder(r.Body).Decode(param)
	if result, err := marketcode.TicketToCode(param); err == nil {
		response.NewResponse(w).OK(result)
	}
}
