package manage_player

import (
	"github.com/name5566/leaf/util"
)

var MPlayer *PlayerManager

func init() {
	MPlayer = NewPlayerManager()
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		mPlayer: new(util.Map),
	}
}

type PlayerManager struct {
	mPlayer *util.Map
}

func (m *PlayerManager) AddPlayer(id int64, p *Player) {
	m.mPlayer.Set(id, p)
}

func (m *PlayerManager) DelPlayer(id int64) {
	m.mPlayer.Del(id)
}

func (m *PlayerManager) GetPlayerById(id int64) *Player {
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

func (m *PlayerManager) GetPlayers() (ps []*Player) {
	m.mPlayer.RLockRange(func(k, v interface{}) {
		ps = append(ps, v.(*Player))
	})
	return
}

func (m *PlayerManager) Free() {
	m.mPlayer.LockRange(func(k, v interface{}) {
		m.mPlayer.UnsafeDel(k)
	})
}
