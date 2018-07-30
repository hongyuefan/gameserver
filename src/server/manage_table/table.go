package manage_table

import (
	js "encoding/json"
	"errors"
	agent "server/manage_agent"
	mc "server/manage_card"
	mp "server/manage_player"
	"server/msg"
	"sync/atomic"
	"time"
)

type Table struct {
	TableId   int64
	Max       int
	Status    int32
	TimeOut   int
	Players   *mp.PlayerManager
	chanCard  chan *mc.Card
	chanClose chan bool
	Winer     *mp.Player
}

func NewTable(tableId int64, max int) *Table {
	return &Table{
		TableId:   tableId,
		Max:       max,
		Players:   mp.NewPlayerManager(),
		chanCard:  make(chan *mc.Card, 2),
		chanClose: make(chan bool, 1),
	}
}

func (m *Table) PlayCard(c *mc.Card) {
	select {
	case m.chanCard <- c:
	default:
	}
}

func (m *Table) TableJoin(p *mp.Player) error {
	if m.Players.Count() > m.Max {
		return errors.New("table full")
	} else if m.Players.Count() < m.Max {
		m.Players.AddPlayer(p.Id, p)
		return nil
	}
	m.Start()
	return nil
}

func (m *Table) TableExit(pId int64) {
	m.Players.DelPlayer(pId)
}

func (m *Table) GetStatus() int32 {
	return atomic.LoadInt32(&m.Status)
}

func (m *Table) TableBroadcast(isSuccess bool, buss msg.BussTypeId, data interface{}) {
	byt, _ := js.Marshal(data)
	agent.MAgent.AgentMC(&msg.Response{Success: isSuccess, BussId: buss, Data: byt}, m.Players.GetPlayersId())
}

func (m *Table) TableBroadcaseExcept(msg interface{}, ids []int64) {
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

func (m *Table) TableMulticast(msg interface{}, ids []int64) {
	agent.MAgent.AgentMC(msg, ids)
}

func (m *Table) TableP2P(msg interface{}, to int64) {
	agent.MAgent.AgentP2P(msg, to)
}

func (m *Table) Start() {
	if atomic.LoadInt32(&m.Status) == msg.Flag_Game_Start {
		return
	}
	m.Status = atomic.AddInt32(&m.Status, 1)
}

func (m *Table) Over() {
	if atomic.LoadInt32(&m.Status) == msg.Flag_Game_Over {
		return
	}
	m.Status = atomic.AddInt32(&m.Status, -1)

	m.Players.Free()

	close(m.chanClose)
}

func (m *Table) f() {

	chanT := time.NewTicker(time.Second * time.Duration(m.TimeOut))
	defer chanT.Stop()

	for {
		select {
		case c := <-m.chanCard:
			m.Players.GetPlayerById(c.PlayerId).AddPlayCard(c)
			if m.JudgeIsReady() {
				m.PK()
			}
		case <-chanT.C:
			if !m.JudgeIsReady() {
				for _, p := range m.Players.GetPlayers() {
					if p.GetPlayCard() == nil {
						c := p.GetCardRand()
						p.AddPlayCard(c)
					}
				}
			}
			m.PK()
		case <-m.chanClose:
			return
		}
	}
}

func (m *Table) PK() {

	players := m.Players.GetPlayers()

	result, _ := players[0].GetPlayCard().Duel(players[1].GetPlayCard())

	switch result {
	case 1:
		m.SetWiner(players[0])
	case 0:
		m.SetWiner(nil)
	case -1:
		m.SetWiner(players[1])
	}
	return
}

func (m *Table) SetWiner(player *mp.Player) {
	m.Winer = player
}

func (m *Table) JudgeIsReady() bool {

	var count int

	players := m.Players.GetPlayers()
	if len(players) != m.Max {
		return false
	}
	for _, player := range players {
		if player.GetPlayCard() != nil {
			count++
		}
	}
	if count != m.Max {
		return false
	}
	return true
}
