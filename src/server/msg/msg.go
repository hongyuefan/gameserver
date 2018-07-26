package msg

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&SendIdentifyCode{})
	Processor.Register(&Response{})
}

type SendIdentifyCode struct {
	Nation        string
	MobileOrEmail string
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

func FailedHandler(agent gate.Agent, buss BussTypeId, err error) {
	agent.WriteMsg(&Response{
		Success: false,
		BussId:  buss,
		Message: err.Error(),
	})
}
