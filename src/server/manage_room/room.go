package manage_room

import (
	agent "server/manage_agent"
	mp "server/manage_player"
	mt "server/manage_table"
	"server/msg"
	"sync/atomic"
	"time"
)

type GameRoom struct {
	GameClassId int64
	CreatorId   int64
	CreatorName string
	RoomId      int64
	RoomPass    string
	CreateTime  int64
	MaxNum      int
	State       int32
	Players     *mp.PlayerManager
	Tables      *mt.TableManager
}

func NewGameRoom(classId, roomId, creatorId int64, max int, creatorName, roomPass string) *GameRoom {
	return &GameRoom{
		GameClassId: classId,
		MaxNum:      max,
		CreatorId:   creatorId,
		CreateTime:  time.Now().Unix(),
		RoomPass:    roomPass,
		RoomId:      roomId,
		State:       0,
		Players:     mp.NewPlayerManager(),
		Tables:      mt.NewTableManager(),
	}
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

func (m *GameRoom) RoomBroadcaseExcept(msg interface{}, ids []int64) {
	aIds := m.Players.GetPlayersId()
	for _, v := range ids {
		for ko, vo := range aIds {
			if v != vo {
				aIds = append(aIds[:ko], aIds[ko+1:]...)
			}
		}
	}
	agent.MAgent.AgentMC(msg, aIds)
}

func (m *GameRoom) RoomMulticast(msg interface{}, ids []int64) {
	agent.MAgent.AgentMC(msg, ids)
}

func (m *GameRoom) RoomGetCards() (nJan, nKen, nPo uint32) {
	pls := m.Players.GetPlayers()
	for _, p := range pls {
		j, k, p := p.GetCards()
		nJan += j
		nKen += k
		nPo += p
	}
	return
}

func (m *GameRoom) RoomP2P(msg interface{}, to int64) {
	agent.MAgent.AgentP2P(msg, to)
}

func (m *GameRoom) GetStatus() int32 {
	return atomic.LoadInt32(&m.State)
}

func (m *GameRoom) Start() {
	if atomic.LoadInt32(&m.State) == msg.Flag_Game_Start {
		return
	}
	m.State = atomic.AddInt32(&m.State, 1)
}

func (m *GameRoom) Over() {
	if atomic.LoadInt32(&m.State) == msg.Flag_Game_Over {
		return
	}
	m.State = atomic.AddInt32(&m.State, -1)
	m.Players.Free()
	m.Tables.Free()
}
