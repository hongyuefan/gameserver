package internal

import (
	ms "server/manage_class"
	mr "server/manage_room"
	"server/msg"
	am "server/util/arithmetic"
	"time"
)

func init() {
	handler(&msg.GameRoomAdd{}, handlerGameRoomAdd)
}

func handlerGameRoomAdd(args []interface{}) {	
	var (
		err error
	)
	
	m := args[0].(*msg.GameRoomAdd)

	r := &mr.GameRoom{
		GameClassId: m.GameClassId,
		CreatorId:   m.CreatorId,
		CreateTime:  time.Now().Unix(),
		RoomPass:    m.RoomPass,
		RoomName:    am.GenCode(4),
	}
	
	if err ms.MClass.GetClassById(m.GameClassId).Rooms.AddRoom(r.RoomName)

}
