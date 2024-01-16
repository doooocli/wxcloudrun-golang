package response

import (
	"encoding/json"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// Response 返回结构
type Response struct {
	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

func (r *Response) OK(data interface{}) {
	r.Send(0, "success", data)
}

func (r *Response) Send(code int, msg string, data interface{}) {

	res := &JsonResult{
		Code:     code,
		ErrorMsg: msg,
		Data:     data,
	}
	content, err := json.Marshal(res)
	if err != nil {
		return
	}
	r.w.Header().Set("content-type", "application/json")
	r.w.Write(content)
}
