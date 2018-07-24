package manage_room

import (
	"errors"
	"fmt"
	agent "server/manage_agent"
	"server/msg"

	"github.com/name5566/leaf/util"
)

var MRoom *RoomManager

func init() {
	MRoom = &RoomManager{
		mRoom: new(util.Map),
	}
}

type RoomManager struct {
	mRoom *util.Map
}

func (m *RoomManager) CreateRoom(name string, rom *msg.RoomCreate) (err error) {
	if m.mRoom.Get(name) != nil {
		err = errors.New(fmt.Sprintf("Room %v already exist", name))
		return
	}
	m.mRoom.Set(name, rom)
	return
}

func (m *RoomManager) RemoveRoom(name string) {
	m.mRoom.Del(name)
}

func (m *RoomManager) GetRoom(name string) *msg.RoomCreate {
	rom := m.mRoom.Get(name)
	if rom != nil {
		return rom.(*msg.RoomCreate)
	}
	return nil
}

func (m *RoomManager) AddPlayer(name string, playerId int64) error {
	rom := m.GetRoom(name)
	if rom == nil {
		return errors.New(fmt.Sprintf("Room %v not exist", name))
	}
	if rom.GetCount() >= rom.GetMax() {
		return errors.New("Room full")
	}
	rom.AddPlayer(playerId)
	return nil
}

func (m *RoomManager) DelPlayer(name string, playerId int64) {
	rom := m.GetRoom(name)
	if rom == nil {
		return errors.New(fmt.Sprintf("Room %v not exist", name))
	}
	rom.DelPlayer(playerId)
	if rom.GetCount() == 0 {
		m.RemoveRoom(name)
	}
	return
}

func (m *RoomManager) RoomBroadcast(name string, msg interface{}) {
	rom := m.GetRoom(name)
	if rom != nil {
		agent.MAgent.AgentMC(msg, rom.GetPlayers())
	}
}

func (m *RoomManager) RoomBroadcaseExcept(name string, msg interface{}, ids []int64) {
	rom := m.GetRoom(name)
	if rom != nil {
		aIds := rom.GetPlayers()
		for _, v := range ids {
			for ko, vo := range aIds {
				if v == vo {
					aIds = append(aIds[:ko], aIds[ko:])
				}
			}
		}
		agent.MAgent.AgentMC(msg, aIds)
	}
}

func (m *RoomManager) RoomP2P(msg interface{}, to int64) {
	agent.MAgent.AgentP2P(msg, to)
}
