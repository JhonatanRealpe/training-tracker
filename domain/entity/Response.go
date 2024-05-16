package entity

import (
	"fmt"
	"github.com/JhonatanRealpe/training-tracker/util"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"messagge,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r *Response) SetMessageWithError(status int, err error, message string) *Response {
	if status >= http.StatusBadRequest && status <= http.StatusNetworkAuthenticationRequired {
		r.Status = status
		r.Message = message
		if util.IsEmptyString(message) {
			statusName := http.StatusText(status)
			r.Message = fmt.Sprintf("error: %v %v", statusName, err.Error())
		}
	}
	return r
}
func (r *Response) SetData(status int, messagge string, data interface{}) {
	r.Status = status
	r.Message = messagge
	r.Data = data
}
