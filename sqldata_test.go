package sqldata

import (
	"testing"
	"github.com/BurntSushi/toml"
	"flag"
	"context"
)

func TestFetchMap(t *testing.T) {

	var sysconfig Config
	var configPath string
	flag.StringVar(&configPath, "config", "my.conf", "server config.")
	flag.Parse()

	configPath = "my.conf"

	if _, err := toml.DecodeFile(configPath, &sysconfig); err != nil {
		t.Fatalf("decode err:%v", err)
	}

	mytest,err := newConnDb(&sysconfig.Db)
	if err!=nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	ctx:=context.Background()
	rows,err := mytest.results(ctx,"SELECT * FROM infos limit 3")
	if err!=nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	datas := mytest.fetchMap(rows)
	t.Logf("gat data : %v",datas)
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})
}
