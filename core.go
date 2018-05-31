package sqldata

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //must import
	"log"
)

//connDb  connect db information
type connDb struct {
	dns        string
	retry      int32
	driverName string
	coreDb     *sql.DB
}

//newConnDb star create a db driver
func newConnDb(conf *configDb) (re *connDb, err error) {
	re = &connDb{
		dns:        conf.Addr,
		retry:      conf.Retry,
		driverName: conf.DriverName,
	}
	re.coreDb, err = re.connect()
	return
}

//connect open the driver
func (cn *connDb) connect() (db *sql.DB, err error) {
	if cn.coreDb != nil {
		return cn.coreDb, nil
	}

	db, err = sql.Open(cn.driverName, cn.dns)
	if err != nil {
		log.Fatalf("connect fail;%v", err)
	} else {
		cn.coreDb = db
	}
	return
}

//GetConnDb get the DB infors
func (cn *connDb) getConnDb() (db *sql.DB) {
	db = cn.coreDb
	return
}

//begin db trans start
func (cn *connDb) begin() (btx *sql.Tx, err error) {
	btx, err = cn.coreDb.Begin()
	return
}

//commit db trans commit
func (cn *connDb) commit(btx *sql.Tx) error {
	return btx.Commit()
}

//rollback db trans reset
func (cn *connDb) rollback(btx *sql.Tx) error {
	return btx.Rollback()
}

//prepare
func (cn *connDb) prepare(query string) (stmt *sql.Stmt, err error) {
	stmt, err = cn.coreDb.Prepare(query)
	return
}

//execFrStmt  operate by exec from stmt
func (cn *connDb) execFromStmt(stmt *sql.Stmt, args ...interface{}) (rs sql.Result, sterr error) {
	rs, sterr = stmt.Exec(args...)
	return
}

//execTx operate by exec from tx
func (cn *connDb) execFromTx(btx *sql.Tx, query string, args ...interface{}) (rs sql.Result, exerr error) {
	rs, exerr = btx.Exec(query, args...)
	return
}

//exec operate by exec from db,a normal op
func (cn *connDb) execFromDb(query string, args ...interface{}) (result sql.Result, err error) {
	result, err = cn.coreDb.Exec(query, args...)
	return
}

//getLastId get the result last id
func (cn *connDb) getLastID(result sql.Result) (num int64, err error) {
	num, err = result.LastInsertId()
	return
}

//rowsAffected result is ok or fail
func (cn *connDb) rowsAffected(result sql.Result) (num int64, err error) {
	num, err = result.RowsAffected()
	return
}

//close close db
func (cn *connDb) close() error {
	return cn.coreDb.Close()
}

//results from sql and args,get the datas
func (cn *connDb) query(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = cn.coreDb.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("sql:%v,result:%v", query, err)
	}
	return
}

//fetchMap reset new datas,become to map type
//return data the same array
func (cn *connDb) fetchMap(rows *sql.Rows) (results resultData) {
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
			log.Fatal("sqldata_fetchMap||err=%v", err)
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
