package sqldata

import (
	"context"
	"testing"
)

type mytestperson struct {
	Id   uint8  `sql:"id"`
	Age  uint8  `sql:"age"`
	Name string `sql:"name"`
}

func TestConnect(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName: "mysql",
		Addr:       "xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8",
		Retry:      2,
	}
	mytest, err := newConnDb(conf)
	if err != nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	ctx := context.Background()
	rows, err := mytest.query(ctx, "SELECT * FROM infos limit 3")
	if err != nil {
		t.Fatalf("===== Connect get data. [err:%v]", err)
	}
	datas := mytest.rowsMap(rows)
	t.Logf("=====Connect gat data : %v", datas)

	//arows, err := mytest.query(ctx, "SELECT id,age,name FROM infos limit 3")
	//if err != nil {
	//	t.Fatalf("===== Connect get data. [err:%v]", err)
	//}

	//myp := myperson{}
	//adatas := mytest.rowsObject(arows,myp)  //@todo here是问题
	//t.Fatalf("=====Connect gat object : %v", adatas)
	//for ppi:=0;ppi<len(adatas);ppi++ {
	//	pppcc:=datas[ppi]
	//	t.Log("======= indata:",pppcc["id"])
	//}
}

func TestPrepare(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName: "mysql",
		Addr:       "xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8",
		Retry:      2,
	}
	mytest, err := newConnDb(conf)
	if err != nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	//ctx := context.Background()
	stmt, err := mytest.prepare("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)")
	defer stmt.Close()
	if err != nil {
		t.Fatalf("prepare. [err:%v]", err)
	}
	result, err := mytest.execFromStmt(stmt, "肖2", 30, "肖2", 30)
	if err != nil {
		t.Fatalf("execFromStmt . [err:%v]", err)
	}
	id, _ := result.LastInsertId()
	t.Logf("insert id:%v", id)

}

func TestCommit(t *testing.T) {
	var conf *configDb
	conf = &configDb{
		DriverName: "mysql",
		Addr:       "xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8",
		Retry:      2,
	}
	mytest, err := newConnDb(conf)
	if err != nil {
		t.Fatalf("fail to connect. [err:%v]", err)
	}
	btx, err := mytest.begin()
	if err != nil {
		t.Fatalf("====begin -> begin . [err:%v]", err)
	}
	//ctx := context.Background()
	stmt, err := mytest.txPrepare(btx, "INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)")
	defer stmt.Close()
	if err != nil {
		t.Fatalf("====prepare from begin. [err:%v]", err)
	}
	result, err := mytest.execFromStmt(stmt, "肖cc2", 30, "肖cc2", 30)
	if err != nil {
		t.Fatalf("execFromStmt . [err:%v]", err)
	}
	id, _ := result.LastInsertId()

	err = mytest.commit(btx)
	if err != nil {
		rerr := mytest.rollback(btx)
		t.Logf("==== rollback :%v", rerr)
		t.Fatalf("====commit -> commit . [err:%v]", err)
	}
	t.Logf("==== commit insert id:%v", id)

}
