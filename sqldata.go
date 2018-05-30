package sqldata

import (
	"context"
	"log"
)

//ResultData
type resultData map[int]map[string]string

//SqlData
type SqlData interface {
	// fetch data information
	MysqlFetchMap(sql string, args ...interface{}) (data resultData, err error)
	// insert data information
	PrepareInsert(sql string, args ...interface{}) (lastId int64, err error)
	// op database information
	PrepareOpAffected(sql string, args ...interface{}) (affectedId int64, err error)
	// Insert
	Insert(sql string, args ...interface{}) (lastId int64, err error)
	// OpAffected
	OpAffected(sql string, args ...interface{}) (affectedId int64, err error)
}

//implSqlData
type implSqlData struct {
	ctx    context.Context //ctx from config
	conndb *connDb
}

//GetDb
func (sd *implSqlData) GetDb() (db *connDb) {
	db = sd.conndb
	return
}

//MysqlFetchMap
func (sd *implSqlData) MysqlFetchMap(sql string, args ...interface{}) (data resultData, err error) {
	conDatabase := sd.conndb
	resultRows, err := conDatabase.query(sd.ctx, sql, args...)
	if err != nil {
		log.Fatal("sqldata_MysqlFetchMap||sql=%s||err=%v", sql, err)
	}
	data = conDatabase.fetchMap(resultRows)
	return
}

//PrepareInsert
func (sd *implSqlData) PrepareInsert(sql string, args ...interface{}) (lastId int64, err error) {
	conDatabase := sd.conndb
	stmt, err := conDatabase.prepare(sql)
	defer stmt.Close()
	result, err := conDatabase.execFromStmt(stmt, args...)
	if err != nil {
		log.Fatal("sqldata_PrepareInsert||sql=%s||err=%v", sql, err)
	}
	lastId, _ = result.LastInsertId()
	return
}

//PrepareOpAffected
func (sd *implSqlData) PrepareOpAffected(sql string, args ...interface{}) (affectedId int64, err error) {
	conDatabase := sd.conndb
	stmt, err := conDatabase.prepare(sql)
	defer stmt.Close()
	result, err := conDatabase.execFromStmt(stmt, args...)
	if err != nil {
		log.Fatal("sqldata_PrepareOpAffected||sql=%s||err=%v", sql, err)
	}
	affectedId, _ = result.RowsAffected()
	return
}

//PrepareInsert
func (sd *implSqlData) Insert(sql string, args ...interface{}) (lastId int64, err error) {
	conDatabase := sd.conndb
	result, err := conDatabase.execFromDb(sql, args...)
	if err != nil {
		log.Fatal("sqldata_Insert||sql=%s||err=%v", sql, err)
	}
	lastId, _ = result.LastInsertId()
	return
}

//PrepareOpAffected
func (sd *implSqlData) OpAffected(sql string, args ...interface{}) (affectedId int64, err error) {
	conDatabase := sd.conndb
	result, err := conDatabase.execFromDb(sql, args...)
	if err != nil {
		log.Fatal("sqldata_OpAffected||sql=%s||err=%v", sql, err)
	}
	affectedId, _ = result.RowsAffected()
	return
}
