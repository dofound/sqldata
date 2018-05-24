package sqldata

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //must import
	"context"
)

//connDb
type connDb struct {
	dns string
	retry int32
	driverName string
	coreDb *sql.DB
}

//newConnDb
func newConnDb(conf *configDb) (re *connDb,err error) {
	re = &connDb{
		dns:conf.Addr,
		retry:conf.Retry,
		driverName:conf.DriverName,
	}
	re.coreDb,err = re.connect()
	return
}

//connect
func (cn *connDb)connect() (db *sql.DB,err error) {
	if cn.coreDb!=nil {
		return cn.coreDb,nil
	} else {
		db, err = sql.Open(cn.driverName, cn.dns)
		if err != nil {
			log.Fatalf("connect fail;%v", err)
		} else {
			cn.coreDb = db
		}
	}
	return
}

//GetConnDb
func (cn *connDb)getConnDb() (db *sql.DB){
	db = cn.coreDb
	return
}

//begin
func (cn *connDb)begin() (btx *sql.Tx,err error) {
	btx,err = cn.coreDb.Begin()
	return
}

//commit
func (cn *connDb)commit(btx *sql.Tx) error {
	return btx.Commit()
}

//rollback
func (cn *connDb)rollback(btx *sql.Tx) error {
	return btx.Rollback()
}

//prepare
func (cn *connDb)prepare(query string) (stmt *sql.Stmt,err error) {
	stmt,err = cn.coreDb.Prepare(query)
	return
}

//execFrStmt
func (cn *connDb)execFromStmt(stmt *sql.Stmt,args ...interface{})(rs sql.Result,sterr error) {
	rs,sterr = stmt.Exec(args)
	return
}

//execTx
func (cn *connDb)execFromTx(btx *sql.Tx,query string, args ...interface{})(rs sql.Result,exerr error) {
	rs,exerr = btx.Exec(query, args)
	return
}

//exec
func (cn *connDb)execFromDb(query string, args ...interface{})(result sql.Result,err error) {
	result,err = cn.coreDb.Exec(query, args)
	return
}

//getLastId
func (cn *connDb)getLastId(result sql.Result) (num int64,err error) {
	num,err = result.LastInsertId()
	return
}

//rowsAffected
func (cn *connDb)rowsAffected(result sql.Result) (num int64,err error) {
	num,err = result.RowsAffected()
	return
}

//results 组装sql信息，对信息进行处理
func (cn *connDb)query(ctx context.Context,query string, args...interface{}) (rows *sql.Rows,err error){
	rows,err = cn.coreDb.QueryContext(ctx,query,args...)
	if err!=nil {
		log.Printf("sql:%v,result:%v",query,err)
	}
	return
}

//fetchMap 获取数据，对数据进行转化成map
//返回的数据是对 数据表字段为key
func (cn *connDb)fetchMap(rows *sql.Rows) (results ResultData) {
	columns, _ := rows.Columns()
	values := make([][]byte, len(columns)) //make a byte slice
	fields := make([]interface{}, len(columns))
	for i := range values {
		fields[i] = &values[i]
	}
	results = make(map[int]map[string]string)
	var ii int
	for rows.Next() {
		if err := rows.Scan(fields...); err != nil {
			log.Printf("rows scan:%v",err)
			continue
			//return
		}
		row := make(map[string]string) //every on line
		for k, v := range values {
			key := columns[k]
			row[key] = string(v) //string info
		}
		results[ii] = row
		ii++
	}
	return
}

