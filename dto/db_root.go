package dto

import "database/sql"

func QueryRDBs(db *sql.DB) ([]string, error) {
	sqlStr := "show databases;"
	var s string
	var ss []string
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func QueryRTables(db *sql.DB, dbName string) ([]string, error) {
	sqlStr := "select table_name" +
		"from information_schema.tables" +
		"swhere table_schema='" + dbName + "' and table_type='base table';"
	var s string
	ss := []string{}
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	return ss, nil
}

func QueryRColumns(db *sql.DB, tableName string) ([]string, error) {
	sqlStr := "select * from " + tableName + " limit 1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	return rows.Columns()
}
