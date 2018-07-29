package manage_room

import (
	"errors"
	"fmt"

	"github.com/name5566/leaf/util"
)

type RoomManager struct {
	mRoom *util.Map
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		mRoom: new(util.Map),
	}
}

func (m *RoomManager) AddRoom(roomId int64, rom *GameRoom) (err error) {
	if m.mRoom.Get(roomId) != nil {
		err = errors.New(fmt.Sprintf("Room %v already exist", roomId))
		return
	}
	m.mRoom.Set(roomId, rom)
	return
}

func (m *RoomManager) DelRoom(roomId int64) {
	m.mRoom.Del(roomId)
}

func (m *RoomManager) GetRoomById(roomId int64) *GameRoom {
	rom := m.mRoom.Get(roomId)
	if rom != nil {
		return rom.(*GameRoom)
	}
	return nil
}

func (m *RoomManager) GetRooms() (roms []*GameRoom) {
	m.mRoom.RLockRange(func(k interface{}, v interface{}) {
		roms = append(roms, v.(*GameRoom))
	})
	return
}

func (m *RoomManager) GetRoomsCount() int {
	return m.mRoom.Len()
}
