package dbServer

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"common"
	"fmt"
	"log"
)

type MysqlApi struct {
	db     *sql.DB
}

func CreateMysqlApi() MysqlApi {
	common.GetConfig()
	mysqlApi := MysqlApi{config: config}
	mysqlApi.connect()
	return mysqlApi
}
func (mysqlApi *MysqlApi) connect() error {
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysqlApi.config.DB.Username,
		mysqlApi.config.DB.Password, mysqlApi.config.DB.Server, mysqlApi.config.DB.Port, mysqlApi.config.DB.Db)
	db, err := sql.Open("mysql", mysqlUrl)
	db.SetMaxOpenConns(mysqlApi.config.DB.PoolNumber)
	mysqlApi.db = db
	return err
}
func (mysqlApi MysqlApi) GetWxApp(appId string) (string, error) {
	stmt, err := mysqlApi.db.Prepare("SELECT appsecret FROM jmqjwxapp WHERE appid = ?")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(appId)
	var appSec string
	for rows.Next() {
		if err := rows.Scan(&appSec); err != nil {
			log.Fatal(err)
		}
	}
	return appSec, err
}
