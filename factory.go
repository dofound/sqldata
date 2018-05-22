package sqldata

import "context"

type Factory struct {
	db configDb `toml:"database"`
	mock SqlData
}

func NewFactory(config *Config) *Factory {
	return &Factory{
		db:config.Db,
	}
}

func (f *Factory)New(ctx context.Context) (re SqlData) {
	if f.mock!=nil {
		re = f.mock
	} else {
		re = &implSqlData{
			ctx:ctx,
			conndb:newConnDb(f.db),
		}
	}
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

}
