package internal

import (
	"reflect"
	mc "server/manage_class"
	"server/msg"

	"github.com/name5566/leaf/gate"
	//	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.GameClassGet{}, handlerGetGameClass)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerGetGameClass(args []interface{}) {
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameClassGet_Code, mc.MClass.GetClass())
}
