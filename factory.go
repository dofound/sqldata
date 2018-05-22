package sqldata

import "context"

//Factory
type Factory struct {
	db configDb `toml:"database"`
	mock SqlData
}

//NewFactory
func NewFactory(config *Config) *Factory {
	return &Factory{
		db:config.Db,
	}
}

//New
func (f *Factory)New(ctx context.Context) (re SqlData) {
	if f.mock!=nil {
		re = f.mock
	} else {
		tmpConndb,_:=newConnDb(&f.db)
		re = &implSqlData{
			ctx:ctx,
			conndb:tmpConndb,
		}
	}
	return
}

//SetMock 设置mock
func (f *Factory) SetMock(mock SqlData) {
	f.mock = mock
}

//ResetMock 重置mock
func (f *Factory) ResetMock() {
	f.SetMock(nil)
}

//Ping 检查参数
func (f *Factory) Ping() (err error) {
	return
}
