package internal

import (
	"errors"
	mp "server/manage_player"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.PlayerGetReq{}, handlerGamePlayerGet)
}

func handlerGamePlayerGet(args []interface{}) {

	m := args[0].(*msg.PlayerGetReq)

	p := mp.MPlayer.GetPlayerById(m.PlayerId)
	if p == nil {
		msg.FailedHandler(args[1].(gate.Agent), msg.Buss_GamePlayerGet_Code, errors.New("Player Offline"))
		return
	}

	rsp := msg.PlayerGetRsp{
		Stars: p.GetStars(),
		Gold:  p.GetGold(),
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GamePlayerGet_Code, rsp)
}
