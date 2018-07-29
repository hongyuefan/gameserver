package manage_table

import (
	"github.com/name5566/leaf/util"
)

type TableManager struct {
	mTable *util.Map
}

func NewTableManager() *TableManager {
	return &TableManager{
		mTable: new(util.Map),
	}
}

func (m *TableManager) AddTable(tId int64, t *Table) {
	m.mTable.Set(tId, t)
}

func (m *TableManager) GetTableById(tId int64) *Table {
	tb := m.mTable.Get(tId)
	if tb != nil {
		return tb.(*Table)
	}
	return nil
}

func (m *TableManager) DelTable(tId int64) {
	m.mTable.Del(tId)
}

func (m *TableManager) GetTables() (tbs []*Table) {
	m.mTable.RLockRange(func(k, v interface{}) {
		tbs = append(tbs, v.(*Table))
	})
	return
}

func (m *TableManager) Free() {
	m.mTable.LockRange(func(k, v interface{}) {
		v.(*Table).Over()
		m.mTable.UnsafeDel(k)
	})
}
