package sqldata

import (
	"testing"
)

func TestConnect(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName:"mysql",
		Addr:"1xxx",
		Retry:2,
	}
	mytest,err := newConnDb(conf)
	if err!=nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	rows,err := mytest.results("SELECT * FROM xmy limit 3")
	if err!=nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	datas := mytest.fetchMap(rows)
	t.Fatalf("gat data : %v",datas)
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})
}
