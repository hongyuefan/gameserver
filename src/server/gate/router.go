package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.UserRegist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.SendIdentifyCode{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.UserLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.GameClass{}, game.ChanRPC)
}
