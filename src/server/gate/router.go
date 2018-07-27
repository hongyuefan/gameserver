package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.PlayerRegist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.SendIdentifyCode{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.PlayerLogin{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.GameClassGetReq{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.GameRoomAdd{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GameRoomGetReq{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GameRoomJoinReq{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GameRoomExitReq{}, game.ChanRPC)
}
