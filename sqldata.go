package sqldata

import (
	"log"
	"context"
)

//ResultData
type ResultData map[int]map[string]string

//SqlData
type SqlData interface {
	// fetch data information
	MysqlFetchMap(sql string,args ...interface{}) (data ResultData,err error)
}

//implSqlData
type implSqlData struct{
	ctx  context.Context //ctx from config
	conndb *connDb
}

//FetchMap
func (sd *implSqlData)MysqlFetchMap(sql string,args ...interface{}) (data ResultData,err error) {
	conDatabase := sd.conndb
	resultRows,err := conDatabase.query(sd.ctx,sql,args...)
	if err!=nil {
		log.Fatalf("fetchquery fail")
	}
	data = conDatabase.fetchMap(resultRows)
	return
}