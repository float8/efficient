package database

type Model struct {
	driverName string
	dbname     string
	key        string
	assColumns map[string]bool
	ass        bool
}

type ModelInterface interface {
	Key() string
	TableName() string
	DriverName() string
	Dbname() string
	Get(name string) interface{}
	Set(name string, value interface{})

	AssColumns() map[string]bool
	AssColumn(column string) bool
	Ass() bool

	UpdateEvent()
	InsertEvent()
	tags() columns
	Ptrs() map[string]interface{}
}

func (m *Model) UpdateEvent() {}

func (m *Model) InsertEvent() {}

func (m *Model) Init(driverName, dbname, key string) *Model {
	m.driverName = driverName
	m.dbname = dbname
	m.key = key
	return m
}

func (m *Model) Key() string {
	return m.key
}

func (m *Model) DriverName() string {
	return m.driverName
}

func (m *Model) Dbname() string {
	return m.dbname
}

func (m *Model) AddAssColumns(column string) {
	if m.assColumns == nil {
		m.assColumns = map[string]bool{}
	}
	if !m.ass {
		m.ass = true
	}
	m.assColumns[column] = true
}

func (m *Model) AssColumns() map[string]bool {
	return m.assColumns
}

func (m *Model) Ass() bool {
	return m.ass
}

func (m *Model) AssColumn(column string) bool {
	if v, ok := m.assColumns[column]; ok {
		return v
	}
	return false
}

func (m *Model) tags() columns {
	return dirs[m.driverName][m.dbname][m.key]
}
