package sqldata

import (
	"log"
	"context"
)

type ResultData map[int]map[string]string

type SqlData interface {
	// fetch data information
	FetchMap(sql string,args ...interface{}) (data *ResultData,err error)
}

type implSqlData struct{
	ctx  context.Context //ctx from config
	conndb *connDb
}

func (sd *implSqlData)FetchMap(sql string,args ...interface{}) (data *ResultData,err error) {
	conDatabase := sd.conndb
	resultRows,err := conDatabase.results(sql,args...)
	if err!=nil {
		log.Fatalf("fetchquery fail")
	}
	data = conDatabase.fetchMap(resultRows)
	return
}