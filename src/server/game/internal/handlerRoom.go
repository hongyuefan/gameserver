package internal

import (
	"time"
	"server/msg"
	mr"server/manage_room"
)

func init() {
	handler(&msg.GameRoomAdd{}, handlerGameRoomAdd)
}

func handlerGameRoomAdd(args []interface{}) {
	
	m := args[0].*msg.GameRoomAdd
	
	r := &mr.GameRoom{
		GameClassId : m.GameClassId,
		CreatorId : m.CreatorId,
		CreateTime : time.Now().Unix(),
		RoomPass: m.RoomPass,
	}
}
