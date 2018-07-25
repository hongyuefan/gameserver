package manage_room

import (
	agent "server/manage_agent"
	mp "server/manage_player"
)

type GameRoom struct {
	GameClassId int64
	CreatorId   int64
	RoomName    string
	RoomPass    string
	CreateTime  int64
	MaxNum      int
	Players     *mp.PlayerManager
}

func (m *GameRoom) Count() int {
	return m.Players.Count()
}

func (m *GameRoom) GetMax() int {
	return m.MaxNum
}

func (m *GameRoom) RoomBroadcast(msg interface{}) {
	agent.MAgent.AgentMC(msg, m.Players.GetPlayersId())
}

func (m *GameRoom) RoomBroadcaseExcept(name string, msg interface{}, ids []int64) {
	aIds := m.Players.GetPlayersId()
	for _, v := range ids {
		for ko, vo := range aIds {
			if v == vo {
				aIds = append(aIds[:ko], aIds[ko+1:]...)
			}
		}
	}
	agent.MAgent.AgentMC(msg, aIds)
}

func (m *GameRoom) RoomMulticast(msg interface{}, ids []int64) {
	agent.MAgent.AgentMC(msg, ids)
}

func (m *GameRoom) RoomP2P(msg interface{}, to int64) {
	agent.MAgent.AgentP2P(msg, to)
}
