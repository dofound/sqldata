package sqldata

import (
	"testing"
	"context"
)

func TestConnect(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName:"mysql",
		Addr:"xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8",
		Retry:2,
	}
	mytest,err := newConnDb(conf)
	if err!=nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	ctx := context.Background()
	rows,err := mytest.query(ctx,"SELECT * FROM infos limit 3")
	if err!=nil {
		t.Fatalf("get data. [err:%v]", err)
	}
	datas := mytest.fetchMap(rows)
	t.Logf("gat data : %v",datas)
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})
}

func TestPrepare(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName:"mysql",
		Addr:"xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8",
		Retry:2,
	}
	mytest,err := newConnDb(conf)
	if err!=nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	//ctx := context.Background()
	stmt,err := mytest.prepare("INSERT INTO `infos` (`name`, `age`) VALUES (?,?)")
	defer stmt.Close()
	if err!=nil {
		t.Fatalf("prepare. [err:%v]", err)
	}
	result,err := mytest.execFromStmt(stmt,"肖2",30)
	//result,err:= stmt.Exec("肖2",30)
	if err!=nil {
		t.Fatalf("execFromStmt . [err:%v]", err)
	}

	t.Log(result.LastInsertId())
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})

}
