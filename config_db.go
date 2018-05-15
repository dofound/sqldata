package sqldata

type configDb struct {
	Addr string `toml:"addr"`
	Retry int32 `toml:"retry"`
	LifeTime int32 `toml:"life_time"`
	DriverName string `toml:"driver_name`
}

func (cf *configDb)addr() string {
	return cf.Addr
}

func (cf *configDb)retry() int32 {
	return cf.Retry
}

func (cf *configDb)lifeTime() int32 {
	return cf.LifeTime
}

func (cf *configDb)driverName() string {
	return cf.DriverName
}