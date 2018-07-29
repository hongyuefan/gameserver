package handler

import (
	"server/msg"
)

type CallFunc func(interface{})

type SendIdentifyCode struct {
	Nation        string
	MobileOrEmail string
}

type UserRegist struct {
	MobileOrEmail string
	Password      string
	VerifyCode    string
}

type UserData struct {
	Id   int64
	Name string
	Msg  string
}

type UserLogin struct {
	MobileOrEmail string
	Password      string
}

type GameClass struct {
	Id              int64
	GameName        string
	GamePlayerCount int64
}

type P struct {
	Id   int64
	Name string
}

type ChatTo struct {
	P
	OpType msg.BussTypeId
	Token  string
	Msg    string
	To_Ids []int64
}

type ChatLogin struct {
	Token string
}

func (c *ChatTo) GetMsg() string {
	return c.Name + ": " + c.Msg
}

type Response struct {
	Success bool
	BussId  msg.BussTypeId
	Message string
	Data    interface{}
}
