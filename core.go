package sqldata

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //must import
)

type connDb struct {
	dns string
	retry int32
	driverName string
	coreDb *sql.DB
}

func (cn *connDb) connect() (db *sql.DB,err error) {
	db,err =sql.Open(cn.driverName,cn.dns)
	if err!=nil {
		log.Fatalf("connect fail;%v",err)
	}
	return
}

func (cn *connDb) results(query string, args ...interface{}) (rows *sql.Rows,err error){
	rows,err = cn.coreDb.Query(query,args...)
	if err!=nil {
		log.Printf("sql:%v,result:%v",query,err)
	}
	return
}

func (cn *connDb) fetchMap(rows *sql.Rows) (results map[int]map[string]string) {
	columns, _ := rows.Columns()
	values := make([][]byte, len(columns)) //make a byte slice
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}
	results = make(map[int]map[string]string)
	var i int32
	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			log.Printf("rows scan:%v",err)
			//continue
			return
		}
		row := make(map[string]string) //every on line
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}
	return
	//for k, v := range results {
	//	fmt.Printf("%d,%v\n",k, v)
	//}
}