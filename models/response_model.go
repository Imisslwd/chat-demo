package models

import "encoding/json"

type Head struct {
	Seq string `json:"seq"`
	Cmd string `json:"cmd"`
	Response *Response `json:"response"`
}

type Response struct {
	
	Code uint `json:"code"`
	CodeMsg string `json:"codeMsg"`
	Data interface{} `json:"data"`
}


func NewResponseHead(seq string, cmd string, code uint, codeMsg string, data interface{}) *Head  {
	response := NewResponse(code, codeMsg, data)
	return &Head{Seq:seq, Cmd: cmd, Response: response}
}

func NewResponse(code uint, codeMsg string, data interface{}) (*Response) {
	
	return &Response{Code: code, CodeMsg: codeMsg, Data: data}
}

func (h *Head)toString() (headStr string) {
	headBytes, _ := json.Marshal(h);
	headStr = string(headBytes)
	return
}

func (h *Head)ToByte() (headByte []byte, err error) {
	headByte, err := json.Marshal(h);
	return
}