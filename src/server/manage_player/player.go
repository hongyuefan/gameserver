package manage_player

import (
	"sync/atomic"

	mc "server/manage_card"
	"server/msg"

	"github.com/name5566/leaf/util"
)

type Player struct {
	Id            int64
	NickName      string
	Stars         uint32
	JKP_Cards     *util.Map
	PlayCard      *mc.Card
	Gold          uint64
	Scole_Success uint32
	Scole_Failed  uint32
	CreateTime    int64
}

func NewPlayer() *Player {
	return &Player{
		JKP_Cards: new(util.Map),
	}
}

func (m *Player) AddStars(x uint32) {
	m.Stars = atomic.AddUint32(&m.Stars, x)
}

func (m *Player) SubStars(x uint32) {
	m.Stars = atomic.AddUint32(&m.Stars, ^uint32(x-1))
}

func (m *Player) GetStars() uint32 {
	return atomic.LoadUint32(&m.Stars)
}

func (m *Player) GetCards() (nJan, nKen, nPo uint32) {
	return m.JKP_Cards.Get(msg.Card_Type_Jan).(uint32), m.JKP_Cards.Get(msg.Card_Type_Ken).(uint32), m.JKP_Cards.Get(msg.Card_Type_Po).(uint32)
}

func (m *Player) GetCardCount() uint32 {
	return m.JKP_Cards.Get(msg.Card_Type_Jan).(uint32) + m.JKP_Cards.Get(msg.Card_Type_Ken).(uint32) + m.JKP_Cards.Get(msg.Card_Type_Po).(uint32)
}

func (m *Player) AddCard(ctp msg.CardTypeId) {
	num := m.JKP_Cards.Get(ctp).(uint32)
	m.JKP_Cards.Set(ctp, num+1)
}

func (m *Player) SubCard(ctp msg.CardTypeId) {
	num := m.JKP_Cards.Get(ctp).(uint32)
	if num <= 0 {
		return
	}
	m.JKP_Cards.Set(ctp, num-1)
}

func (m *Player) GetCardRand() *mc.Card {
	if m.GetCardCount() > 0 {
		m.JKP_Cards.Get()
	}
}

func (m *Player) GetGold() uint64 {
	return atomic.LoadUint64(&m.Gold)
}

func (m *Player) AddGold(x uint64) {
	atomic.AddUint64(&m.Gold, x)
}

func (m *Player) SubGold(x uint64) {
	atomic.AddUint64(&m.Gold, ^uint64(x-1))
}

func (m *Player) AddPlayCard(c *mc.Card) {
	m.PlayCard = c
}

func (m *Player) GetPlayCard() *mc.Card {
	return m.PlayCard
}

func (m *Player) DelPlayCard() {
	m.PlayCard = nil
}
