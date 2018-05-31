package sqldata

import "context"

//Factory define factory config
type Factory struct {
	db   configDb `toml:"database"`
	mock SQLData
}

//NewFactory factory
func NewFactory(config *Config) *Factory {
	return &Factory{
		db: config.Db,
	}
}

//New new implSqlData
func (f *Factory) New(ctx context.Context) (re SQLData) {
	if f.mock != nil {
		re = f.mock
	} else {
		tmpConndb, _ := newConnDb(&f.db)
		re = &implSQLData{
			ctx:    ctx,
			conndb: tmpConndb,
		}
	}
	return
}

//SetMock set mock
func (f *Factory) SetMock(mock SQLData) {
	f.mock = mock
}

//ResetMock set mock
func (f *Factory) ResetMock() {
	f.SetMock(nil)
}

//Ping check
func (f *Factory) Ping() (err error) {
	return
}
