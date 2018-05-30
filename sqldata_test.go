package sqldata

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"testing"
)

func TestMysqlFetchMap(t *testing.T) {

	var sysconfig Config
	var configPath string
	flag.StringVar(&configPath, "config", "my.conf", "server config.")
	flag.Parse()

	configPath = "my.conf"

	if _, err := toml.DecodeFile(configPath, &sysconfig); err != nil {
		t.Fatalf("decode err:%v", err)
	}
	newSql := NewFactory(&sysconfig)
	ctx := context.Background()
	sqlHand := newSql.New(ctx)

	condition := 2
	datas, err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?", condition)
	if err != nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	for pkey, val := range datas {
		t.Log("%v,%v", pkey, val)
	}
	t.Logf("gat data : %v", datas)
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})
}
