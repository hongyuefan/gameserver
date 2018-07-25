package manage_player

import (
	"github.com/name5566/leaf/util"
)

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		mPlayer: new(util.Map),
	}
}

type PlayerManager struct {
	mPlayer *util.Map
}

func (m *PlayerManager) AddtPlayer(id int64, p *Player) {
	m.mPlayer.Set(id, p)
}

func (m *PlayerManager) DelPlayer(id int64) {
	m.mPlayer.Del(id)
}

func (m *PlayerManager) GetPlayer(id int64) *Player {
	p := m.mPlayer.Get(id)
	if p != nil {
		return p.(*Player)
	}
	return nil
}

func (m *PlayerManager) UpdatePlayer(id int64, p *Player) {
	m.mPlayer.Set(id, p)
}

func (m *PlayerManager) Count() int {
	return m.mPlayer.Len()
}

func (m *PlayerManager) GetPlayersId() (ids []int64) {
	m.mPlayer.RLockRange(func(k, v interface{}) {
		ids = append(ids, k.(int64))
	})
	return
}
