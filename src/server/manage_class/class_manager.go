package manage_class

import (
	_ "server/database/mysqlbase"
	db "server/database/mysqlbase"
	"server/msg"

	"github.com/name5566/leaf/util"
)

var MClass *ClassManager

func init() {
	MClass = NewClassManager()
	query := make(map[string]string, 0)
	ml, _, err := db.GetAllGameClass(query, []string{}, []string{"game_id"}, []string{"desc"}, 0, 100)
	if err != nil {
		panic(err)
	}
	for _, v := range ml {
		MClass.AddClass(v.(db.GameClass).Id, v.(db.GameClass).GameName, v.(db.GameClass).GamePlayerCount)
	}
}

type ClassManager struct {
	mClass *util.Map
}

func NewClassManager() *ClassManager {
	return &ClassManager{
		mClass: new(util.Map),
	}
}

func (m *ClassManager) AddClass(id int64, name string, max int) {
	m.mClass.Set(id, NewGameClass(id, name, max))
}

func (m *ClassManager) DelClass(id int64) {
	m.mClass.Del(id)
}

func (m *ClassManager) GetClassById(id int64) *GameClass {
	c := m.mClass.Get(id)
	if c != nil {
		return c.(*GameClass)
	}
	return nil
}

func (m *ClassManager) GetClass() (gcs []*msg.GameClass) {
	m.mClass.RLockRange(func(k, v interface{}) {
		gcs = append(gcs, v.(*msg.GameClass))
	})
	return
}
