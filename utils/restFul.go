package utils

import (
	"github.com/yangguifnag/simplex2go/common"
	"time"
)

type RestFulMsg struct {
	Success  bool        `json:"success"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
	Info     string      `json:"info,omitempty"`
	DateTime string      `json:"dateTime,omitempty"`
	Version  string      `json:"version,omitempty"`
	Type     string      `json:"type,omitempty"`
}

type RestFul struct {
	CodeConfig common.HttpCodeConfig
	RestFulMsg RestFulMsg
}

func (r *RestFul) SetSuccess(status string) *RestFul {
	r.RestFulMsg.Success = true
	code := r.CodeConfig.Success[status]
	if code == 0 {
		code = 200
	}
	r.RestFulMsg.Code = code
	return r
}

func (r *RestFul) SetFail(status string) *RestFul {
	r.RestFulMsg.Success = false
	code := r.CodeConfig.Fail[status]
	if code == 0 {
		code = 500
	}
	r.RestFulMsg.Code = code
	return r
}

func (r *RestFul) SetMessage(message string) *RestFul {
	r.RestFulMsg.Message = message
	return r
}

func (r *RestFul) SetData(data interface{}) *RestFul {
	r.RestFulMsg.Data = data
	return r
}

func (r *RestFul) GetJson() RestFulMsg {
	r.RestFulMsg.DateTime = time.Now().Format("2006-01-02 15:04:05")
	r.RestFulMsg.Version = "1.0"
	r.RestFulMsg.Type = "json"
	return r.RestFulMsg
}

func RestFulJson(data interface{}) RestFulMsg {
	return RestFulMsg{
		Success:  true,
		Code:     200,
		Message:  "ok",
		Data:     data,
		DateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func RestFulJsonFail(err string, msg string) RestFulMsg {
	return RestFulMsg{
		Success:  false,
		Code:     500,
		Info:     err,
		Message:  msg,
		DateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
}
