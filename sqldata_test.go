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
	datas, err := sqlHand.FetchMapFromSql("SELECT * FROM infos where id=?", condition)
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
	row := sqlHand.GetDb().QueryRow("SELECT name FROM infos where id=?", 2)
	var name string
	err := row.Scan(&name)
	if err != nil {
		t.Fatalf("get error. [err:%v]", err)
	}
	t.Logf("TestGetDb op : %v", name)

}

func TestInterfaceCommit(t *testing.T) {
	initConfig()
	sqlHand.Begin()
	stmt, err := sqlHand.TxPrepare("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?")
	result, err := sqlHand.TxExec(stmt, "hahha", 2)

	stmt, err = sqlHand.TxPrepare("INSERT INTO `infos` (`name`, `age`) VALUES (?,?)")
	result, err = sqlHand.TxExec(stmt, "肖commit1", 30)
	lastid, err := result.LastInsertId()
	t.Logf("======InterfaceCommit result 3 op : %v", lastid)

	stmt, err = sqlHand.TxPrepare("INSERT INTO `infos` (`name`, `age`) VALUES (?,?)")
	result, err = sqlHand.TxExec(stmt, 3)
	t.Logf("======InterfaceCommit result 4 op : %v", err)

	err = sqlHand.Commit()
	if err != nil {
		sqlHand.Rollback()
		t.Fatalf("InterfaceCommit error. [err:%v]", err)
	}
	t.Logf("InterfaceCommit op is ok")
}
