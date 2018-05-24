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
func (cn *connDb)GetConnDb() (db *sql.DB){
	db = cn.coreDb
	return
}

//results 组装sql信息，对信息进行处理
//
func (cn *connDb)results(ctx context.Context,query string, args...interface{}) (rows *sql.Rows,err error){
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

//beginOp
func (cn *connDb) beginOp() {

}
