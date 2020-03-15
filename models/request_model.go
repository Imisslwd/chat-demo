package models

type Request struct {
	Seq  string      `json:"seq"`
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data,omitempty"`
}

type Login struct {
	Account      string `json:"account"`
	Password     string `json:"password"`
	ServiceToken string `json:"serviceToken"`
	UserId       string `json:"userId, omitempty"`
	AppId        string `json:"appId, omitempty"`
}

type HeartBeat struct {
	UserId string `json:"userId, omitempty"`
}
