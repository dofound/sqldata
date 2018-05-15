package sqldata

import (
	"log"
)

type ResultData map[int]map[string]string

type SqlData interface {
	// fetch data information
	FetchMap(sql string,args ...interface{}) (data *ResultData,err error)
}

func (sd *SqlData)FetchMap(sql string,args ...interface{}) (data *ResultData,err error) {
	conDatabase := &connDb{

	}
	resultRows,err := conDatabase.results(sql,args...)
	if err!=nil {
		log.Fatalf("fetchquery fail")
	}
	data = conDatabase.fetchMap(resultRows)
	return
}