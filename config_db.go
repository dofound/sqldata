package sqldata

type ConfigDb struct {
	Addr string `toml:"addr"`
	Retry int32 `toml:"retry"`
	LifeTime int32 `toml:"life_time"`
	DriverName string `toml:"driver_name`
}

func (cf *ConfigDb)addr() string {
	return cf.Addr
}

func (cf *ConfigDb)retry() int32 {
	return cf.Retry
}

func (cf *ConfigDb)lifeTime() int32 {
	return cf.LifeTime
}

func (cf *ConfigDb)driverName() string {
	return cf.DriverName
}