package msg

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&GameClass{})
	Processor.Register(&UserRegist{})
	Processor.Register(&UserLogin{})
	Processor.Register(&RoomCreate{})
	Processor.Register(&SendIdentifyCode{})
	Processor.Register(&Response{})
}

type SendIdentifyCode struct {
	Nation        string
	MobileOrEmail string
}

type UserRegist struct {
	MobileOrEmail string
	Password      string
	VerifyCode    string
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

type RoomCreate struct {
	GameClassId int64
	PlayerId    int64
	RoomName    string
	RoomPass    string
	CreateTime  int64
}

type Response struct {
	Success bool
	BussId  BussTypeId
	Message string
	Data    interface{}
}

func SuccessHandler(agent gate.Agent, buss BussTypeId, data interface{}) {
	agent.WriteMsg(&Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	})
}

func FailedHandler(agent gate.Agent, err error) {
	agent.WriteMsg(&Response{
		Success: false,
		Message: err.Error(),
	})
}
