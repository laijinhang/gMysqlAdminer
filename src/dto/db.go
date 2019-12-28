package dto

import "database/sql"

// db info
type DBInfo struct {
	Name string `json:"name"`
	Pwd string `json:"pwd"`
	DBDriver string `json:"driver"`
	Addr string `json:"addr"`
	DBName string `json:"db_name"`
}

// login request
type LoginReq struct {
	DBInfo
}

// login resp
type LoginResp struct {
	Resp
}

type Resp struct {
	Code int `json:"int"`
	Message string `json:"message"`
	Data interface{} `json:"message"`
}

type SqlReq struct {
	Cmd string `json:"cmd"`
}

type SqlResp struct {
	Resp
}

// 查询当前用户的所有数据库
func QueryDBs(db *sql.DB) {

}

// 查询某个数据库下的所有表,会切换当前的数据库
func QueryTables(db *sql.DB, dbName string) ([]string, error) {

	return nil, nil
}

// 查询表下的所有字段
func QueryColumns(db *sql.DB, tableName string) ([]string, error) {
	return QueryRColumns(db, tableName)
}