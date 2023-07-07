package models

import "net/http"

type BaseResponse struct {
	Status  int         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var InvalidParameterResponse BaseResponse = BaseResponse{
	Status:  http.StatusBadRequest,
	Message: "Invalid Parameters",
	Error:   "Invalid Paramerters",
}
var BadGatewayResponse BaseResponse = BaseResponse{
	Status:  http.StatusBadGateway,
	Message: "Something Went wrong",
	Error:   "Somthing went wrong",
}
