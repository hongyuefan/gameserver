package manage_table

import (
	"errors"
	"fmt"

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

func (m *TableManager) AddTable(tId int, t *Table) {
	m.mTable.Set(tId, t)
}

func (m *TableManager) GetTableById(tId int) *Table {
	return m.mTable.Get(tId)
}

func (m *TableManager) DelTable(tId int) {
	m.mTable.Del(tId)
}

func (m *TableManager) GetTables() (tbs []*Table) {
	m.mTable.RLockRange(func(k, v interface{}) {
		tbs = append(tbs, v.(*Table))
	})
}
