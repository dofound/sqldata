package sqldata

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"testing"
)

var sqlHand SQLData

func initConfig() {}

func TestInitConfig(t *testing.T) {

	if sqlHand == nil {
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
		t.Fatalf("get error. [err:%v]", err)
	}
	for pkey, val := range datas {
		t.Log("%v,%v", pkey, val)
	}
	t.Logf("TestMysqlFetchMap : %v", datas)
}

func TestPrepareInsert(t *testing.T) {
	initConfig()
	lastId, err := sqlHand.PrepareInsert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)",
		"肖2", 30, "肖2", 30)
	if err != nil {
		t.Fatalf("get error. [err:%v]", err)
	}
	t.Logf("PrepareInsert insert_id : %v", lastId)

}

func TestPrepareOpAffected(t *testing.T) {
	initConfig()
	affect, err := sqlHand.PrepareOpAffected("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?",
		"xiaojh12", 26)
	if err != nil {
		t.Fatalf("get error. [err:%v]", err)
	}
	t.Logf("PrepareOpAffected op : %v", affect)

}

func TestOpAffected(t *testing.T) {
	initConfig()
	affect, err := sqlHand.OpAffected("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?",
		"xiaojh22", 22)
	if err != nil {
		t.Fatalf("get error. [err:%v]", err)
	}
	t.Logf("OpAffected op : %v", affect)

}

func TestGetDb(t *testing.T) {
	initConfig()
	//affect, err := sqlHand
	//if err != nil {
	//	t.Fatalf("get error. [err:%v]", err)
	//}
	//t.Logf("OpAffected op : %v", affect)

}
