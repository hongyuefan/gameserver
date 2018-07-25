package internal

import (
	ms "server/manage_class"
	mp "server/manage_player"
	mr "server/manage_room"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.GameRoomAdd{}, handlerGameRoomAdd)
	handler(&msg.GameRoomGetReq{}, handlerGameRoomGet)
}

func handlerGameRoomAdd(args []interface{}) {
	var (
		err error
	)
	m := args[0].(*msg.GameRoomAdd)

	gameClass := ms.MClass.GetClassById(m.GameClassId)

	gamePlayer := mp.MPlayer.GetPlayer(m.CreatorId)

	gameRoom := mr.NewGameRoom(m.GameClassId, m.CreatorId, gameClass.GetMax(), gamePlayer.NickName, m.RoomPass)

	gameRoom.Players.AddPlayer(m.CreatorId, gamePlayer)

	if err = gameClass.Rooms.AddRoom(gameRoom.RoomName, gameRoom); err != nil {
		msg.FailedHandler(args[1].(gate.Agent), err)
	}

	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameRoomAdd_Code, nil)

}

func handlerGameRoomGet(args []interface{}) {
	var (
		bflag     bool
		gameRooms []*msg.GameRoomGetRsp
	)
	m := args[0].(*msg.GameRoomGetReq)

	roms := ms.MClass.GetClassById(m.ClassId).Rooms.GetRooms()

	for _, rom := range roms {
		if rom.RoomPass != "" {
			bflag = true
		}
		gameRooms = append(gameRooms, &msg.GameRoomGetRsp{
			RoomName:    rom.RoomName,
			CreatorName: rom.CreatorName,
			IsNeedPass:  bflag,
			Creatime:    rom.CreateTime,
		})
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameRoomGet_Code, gameRooms)
}
