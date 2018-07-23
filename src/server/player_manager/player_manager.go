package player_manager

import (
	"github.com/name5566/leaf/util"
)

var MPlayer *PlayerManager

func init() {
	MPlayer = new(PlayerManager)
}

type PlayerManager struct {
	mPlayer util.Map
}

type Player struct {
	Id         int64
	NickName   string
	CreateTime int64
}

func (m *PlayerManager) InsertPlayer(id int64, nickName string, creatime int64) {
	m.mPlayer.Set(id, &Player{
		Id:         id,
		NickName:   nickName,
		CreateTime: creatime,
	})
}

func (m *PlayerManager) DeletePlayer(id int64) {
	m.mPlayer.Del(id)
}

func (m *PlayerManager) GetPlayer(id int64) *Player {
	return m.mPlayer.Get(id)
}
