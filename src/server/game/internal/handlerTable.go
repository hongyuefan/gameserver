package internal

import (
	ms "server/manage_class"
	mp "server/manage_player"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(msg.GameTableJoinReq{}, handlerGameTableJoin)
	handler(msg.GameTableExitReq{}, handlerGameTableExit)
}

func handlerGameTableJoin(args []interface{}) {

	m := args[0].(*msg.GameTableJoinReq)

	table := ms.MClass.GetClassById(m.ClassId).Rooms.GetRoomById(m.RoomId).Tables.GetTableById(m.TableId)

	err := table.TableJoin(mp.MPlayer.GetPlayerById(m.PlayerId))
	if err != nil {
		msg.FailedHandler(args[1].(gate.Agent), msg.Buss_GameTableJoin_Code, err)
		return
	}
	rsp := &msg.GameTableJoinRsp{
		TableId:  m.TableId,
		PlayerId: m.PlayerId,
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameTableJoin_Code, rsp)
	return
}

func handlerGameTableExit(args []interface{}) {

	m := args[0].(*msg.GameTableExitReq)

	table := ms.MClass.GetClassById(m.ClassId).Rooms.GetRoomById(m.RoomId).Tables.GetTableById(m.TableId)

	table.TableExit(m.PlayerId)

	rsp := &msg.GameTableExitRsp{
		TableId:  m.TableId,
		PlayerId: m.PlayerId,
	}

	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameTableExit_Code, rsp)

	return
}

func handlerGameTablePK(args []interface{}) {

	m := args[0].(*msg.GameTablePKReq)

}
