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

func (m *RoomManager) AddRoom(name string, rom *GameRoom) (err error) {
	if m.mRoom.Get(name) != nil {
		err = errors.New(fmt.Sprintf("Room %v already exist", name))
		return
	}
	m.mRoom.Set(name, rom)
	return
}

func (m *RoomManager) DelRoom(name string) {
	m.mRoom.Del(name)
}

func (m *RoomManager) GetRoomByName(name string) *GameRoom {
	rom := m.mRoom.Get(name)
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
