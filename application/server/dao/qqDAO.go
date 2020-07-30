package dao

import (
	"QQ/application/server/conf"
	"database/sql"
	"errors"
	"fmt"

	// "fmt"
	// "log"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	Db *sql.DB
}

var (
	MyDao *Dao
)

func NewDao() (myDao *Dao) {
	db, err := sql.Open(conf.DirverName, conf.DataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(500)
	db.Ping()
	myDao = &Dao{
		Db: db,
	}
	return
}

func (this *Dao) ExecuteDql(sql string) (res []map[string]string, err error) {

	rows, err := this.Db.Query(sql)
	if err != nil {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	var i int
	for i = 0; rows.Next(); i++ {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		// fmt.Println(record)
		res = append(res, record)
	}
	if i == 0 {
		return res, errors.New("Empty Set")
	}
	return
}

func (this *Dao) ExecuteDqlOne(sql string) (res map[string]string, err error) {
	// fmt.Println(sql)
	rows, err := this.Db.Query(sql)
	if err != nil {
		return
	}

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	res = make(map[string]string)
	if !rows.Next() {
		return res, errors.New("Empty Set")
	}
	err = rows.Scan(scanArgs...)
	for i, col := range values {
		if col != nil {
			res[columns[i]] = string(col.([]byte))
		}
	}
	return
}

func (this *Dao) ExecuteDml(sql string) (rowsAffected int64, err error) {
	res, err := this.Db.Exec(sql)
	if err != nil {
		// fmt.Println(err)
		return
	}
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return
	}
	if rowsAffected == 0 {
		err = errors.New("无效语句")
	}
	return
}
