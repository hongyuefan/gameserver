package msg

import (
	js "encoding/json"

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
	Data    []byte
}

func SuccessHandler(agent gate.Agent, buss BussTypeId, data interface{}) {

	byt, _ := js.Marshal(data)

	agent.WriteMsg(&Response{
		Success: true,
		BussId:  buss,
		Data:    byt,
	})
}

func FailedHandler(agent gate.Agent, buss BussTypeId, err error) {
	agent.WriteMsg(&Response{
		Success: false,
		BussId:  buss,
		Message: err.Error(),
	})
}
