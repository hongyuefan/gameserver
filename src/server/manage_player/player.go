package manage_player

type Player struct {
	Id         int64
	NickName   string
	CreateTime int64
}

func NewPlayer() *Player {
	return &Player{}
}
