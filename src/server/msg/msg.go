package msg

import (
	"sync"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&GameClass{})
	Processor.Register(&UserRegist{})
	Processor.Register(&UserLogin{})
	Processor.Register(&RoomCreate{})
	Processor.Register(&SendIdentifyCode{})
	Processor.Register(&Response{})
}

type SendIdentifyCode struct {
	Nation        string
	MobileOrEmail string
}

type UserRegist struct {
	MobileOrEmail string
	Password      string
	VerifyCode    string
}

type UserLogin struct {
	MobileOrEmail string
	Password      string
}

type GameClass struct {
	Id              int64
	GameName        string
	GamePlayerCount int64
}

type RoomCreate struct {
	GameClassId int64
	PlayerId    int64
	RoomName    string
	RoomPass    string
	CreateTime  int64
	Players     []int64
	MaxNum      int8
	lock        sync.RWMutex
}

func (m *RoomCreate) AddPlayer(playerId int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.Players = append(m.Players, playerId)
}

func (m *RoomCreate) DelPlayer(playerId int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for k, v := range m.Players {
		if v == playerId {
			m.Players = append(m.Players[:k], m.Players[k:])
		}
	}
}
func (m *RoomCreate) GetPlayers() []int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.Players
}
func (m *RoomCreate) GetCount() int8 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.Players)
}

func (m *RoomCreate) GetMax() int8 {
	return m.MaxNum
}

type Response struct {
	Success bool
	BussId  BussTypeId
	Message string
	Data    interface{}
}

func SuccessHandler(agent gate.Agent, buss BussTypeId, data interface{}) {
	agent.WriteMsg(&Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	})
}

func FailedHandler(agent gate.Agent, err error) {
	agent.WriteMsg(&Response{
		Success: false,
		Message: err.Error(),
	})
}
