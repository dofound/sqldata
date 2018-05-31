package sqldata

import (
	"context"
	"log"
	"database/sql"
)

//ResultData is map setting
type resultData map[int]map[string]string

//SQLData interface set
//Define the development method, which will provide better service.
type SQLData interface {
	// Gets native objects, supports access to existing SQL methods
	GetDb() (db *sql.DB)
	// fetch data information
	MysqlFetchMap(sql string, args ...interface{}) (data resultData, err error)
	// insert data information
	PrepareInsert(sql string, args ...interface{}) (lastID int64, err error)
	// op database information
	PrepareOpAffected(sql string, args ...interface{}) (affectedID int64, err error)
	// Insert
	Insert(sql string, args ...interface{}) (lastID int64, err error)
	// OpAffected
	OpAffected(sql string, args ...interface{}) (affectedID int64, err error)
}

//implSqlData impl information
type implSQLData struct {
	ctx    context.Context //ctx from config
	conndb *connDb         //db object
}

//GetDb get CoreDb
func (sd *implSQLData) GetDb() (db *sql.DB) {
	if sd.conndb==nil {
		log.Fatal("sqldata_GetDb||db is null")
	}
	db = sd.conndb.coreDb
	return
}

//MysqlFetchMap get the datas,He's a map format
//The only way to get data is to return the information you want.
func (sd *implSQLData) MysqlFetchMap(sql string, args ...interface{}) (data resultData, err error) {
	conDatabase := sd.conndb
	resultRows, err := conDatabase.query(sd.ctx, sql, args...)
	if err != nil {
		log.Fatal("sqldata_MysqlFetchMap||sql=%s||err=%v", sql, err)
	}
	data = conDatabase.fetchMap(resultRows)
	return
}

//PrepareInsert Batch insert data
func (sd *implSQLData) PrepareInsert(sql string, args ...interface{}) (lastID int64, err error) {
	conDatabase := sd.conndb
	stmt, err := conDatabase.prepare(sql)
	defer stmt.Close()
	result, err := conDatabase.execFromStmt(stmt, args...)
	if err != nil {
		log.Fatal("sqldata_PrepareInsert||sql=%s||err=%v", sql, err)
	}
	lastID, _ = result.LastInsertId()
	return
}

//PrepareOpAffected Batch op data
func (sd *implSQLData) PrepareOpAffected(sql string, args ...interface{}) (affectedID int64, err error) {
	conDatabase := sd.conndb
	stmt, err := conDatabase.prepare(sql)
	defer stmt.Close()
	result, err := conDatabase.execFromStmt(stmt, args...)
	if err != nil {
		log.Fatal("sqldata_PrepareOpAffected||sql=%s||err=%v", sql, err)
	}
	affectedID, _ = result.RowsAffected()
	return
}

//Insert  Single insert data
func (sd *implSQLData) Insert(sql string, args ...interface{}) (lastID int64, err error) {
	conDatabase := sd.conndb
	result, err := conDatabase.execFromDb(sql, args...)
	if err != nil {
		log.Fatal("sqldata_Insert||sql=%s||err=%v", sql, err)
	}
	lastID, _ = result.LastInsertId()
	return
}

//OpAffected Single op data
func (sd *implSQLData) OpAffected(sql string, args ...interface{}) (affectedID int64, err error) {
	conDatabase := sd.conndb
	result, err := conDatabase.execFromDb(sql, args...)
	if err != nil {
		log.Fatal("sqldata_OpAffected||sql=%s||err=%v", sql, err)
	}
	affectedID, _ = result.RowsAffected()
	return
}
