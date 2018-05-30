package sqldata

type configDb struct {
	Addr       string `toml:"addr"`
	Retry      int32  `toml:"retry"`
	LifeTime   int32  `toml:"life_time"`
	DriverName string `toml:"driver_name"`
}

//addr
func (cf *configDb) addr() string {
	return cf.Addr
}

//retry
func (cf *configDb) retry() int32 {
	return cf.Retry
}

//lifeTime
func (cf *configDb) lifeTime() int32 {
	return cf.LifeTime
}

//driverName
func (cf *configDb) driverName() string {
	return cf.DriverName
}
