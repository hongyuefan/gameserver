package internal

import (
	"errors"
	ms "server/manage_class"
	mp "server/manage_player"
	mr "server/manage_room"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func init() {
	handler(&msg.GameRoomAdd{}, handlerGameRoomAdd)
	handler(&msg.GameRoomGetReq{}, handlerGameRoomGet)
	handler(&msg.GameRoomJoinReq{}, handlerGameRoomJoin)
}

func handlerGameRoomAdd(args []interface{}) {
	var (
		err error
	)
	m := args[0].(*msg.GameRoomAdd)

	gameClass := ms.MClass.GetClassById(m.GameClassId)

	gamePlayer := mp.MPlayer.GetPlayerById(m.CreatorId)

	romId := int64(gameClass.Rooms.GetRoomsCount() + 1)

	gameRoom := mr.NewGameRoom(m.GameClassId, romId, m.CreatorId, gameClass.GetMax(), gamePlayer.NickName, m.RoomPass)

	gameRoom.Players.AddPlayer(m.CreatorId, gamePlayer)

	if err = gameClass.Rooms.AddRoom(romId, gameRoom); err != nil {
		msg.FailedHandler(args[1].(gate.Agent), msg.Buss_GameRoomAdd_Code, err)
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
			RoomId:      rom.RoomId,
			CreatorName: rom.CreatorName,
			IsNeedPass:  bflag,
			Creatime:    rom.CreateTime,
		})
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameRoomGet_Code, gameRooms)
}

func handlerGameRoomJoin(args []interface{}) {
	var (
		err        error
		gamePlayer *mp.Player
		rsp        *msg.GameRoomJoinRsp
	)

	m := args[0].(*msg.GameRoomJoinReq)

	rom := ms.MClass.GetClassById(m.ClassId).Rooms.GetRoomById(m.RoomId)

	if rom.Count() >= rom.GetMax() {
		err = errors.New("romm full")
		goto errDeal
	}
	if rom.RoomPass != m.RoomPass {
		err = errors.New("room password not right")
		goto errDeal
	}
	gamePlayer = mp.MPlayer.GetPlayerById(m.PlayerId)
	if gamePlayer == nil {
		err = errors.New("player not exist or offline")
		goto errDeal
	}
	rom.Players.AddPlayer(m.PlayerId, gamePlayer)

	rsp = &msg.GameRoomJoinRsp{
		RoomId:  m.RoomId,
		Players: rom.Players.GetPlayers(),
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameRoomJoin_Code, rsp)
	return
errDeal:
	msg.FailedHandler(args[1].(gate.Agent), msg.Buss_GameRoomJoin_Code, err)
	return

}

func handlerGameRoomExit(args []interface{}) {
	m := args[0].(*msg.GameRoomExitReq)
	rom := ms.MClass.GetClassById(m.ClassId).Rooms.GetRoomById(m.RoomId)
	rom.Players.DelPlayer(m.PlayerId)
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_GameRoomExit_Code, "")
}
