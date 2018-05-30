package sqldata

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"testing"
)

var sqlHand SqlData

func initConfig() {}

func TestInitConfig(t *testing.T) {

	if sqlHand==nil {
		var sysconfig Config
		var configPath string
		flag.StringVar(&configPath, "config", "my.conf", "server config.")
		flag.Parse()

		configPath = "my.conf"

		if _, err := toml.DecodeFile(configPath, &sysconfig); err != nil {
			t.Fatalf("decode err:%v", err)
		}
		t.Log("connetct db is ok")
		newSql := NewFactory(&sysconfig)
		ctx := context.Background()
		sqlHand = newSql.New(ctx)
	}

}

func TestMysqlFetchMap(t *testing.T) {
	initConfig()
	condition := 2
	datas, err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?", condition)
	if err != nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	for pkey, val := range datas {
		t.Log("%v,%v", pkey, val)
	}
	t.Logf("TestMysqlFetchMap : %v", datas)
}

func TestPrepareInsert(t *testing.T) {
	initConfig()
	lastId, err := sqlHand.PrepareInsert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)",
		"肖2", 30,"肖2", 30)
	if err != nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	t.Logf("PrepareInsert insert_id : %v", lastId)

}