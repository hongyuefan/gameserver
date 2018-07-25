package manage_class

import (
	mr "server/manage_room"
)

type GameClass struct {
	Id       int64
	GameName string
	MaxNum   int
	Rooms    *mr.RoomManager
}

func NewGameClass(id int64, name string, max int) *GameClass {
	return &GameClass{
		Id:       id,
		GameName: name,
		MaxNum:   max,
		Rooms:    mr.NewRoomManager(),
	}
}
