package dbServer

import (
	"common"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MysqlApi struct {
	db *sql.DB
}

func CreateMysqlApi() MysqlApi {
	mysqlApi := MysqlApi{}
	mysqlApi.connect()
	return mysqlApi
}
func (mysqlApi *MysqlApi) connect() error {
	config := common.GetConfig()
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", config.DB.Username,
		config.DB.Password, config.DB.Server, config.DB.Port, config.DB.Db)
	db, err := sql.Open("mysql", mysqlUrl)
	db.SetMaxOpenConns(config.DB.PoolNumber)
	mysqlApi.db = db
	return err
}
func (mysqlApi *MysqlApi) GetWxApp(appId string) (string, error) {
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
func (mysqlApi *MysqlApi) SaveOpenId(appId string, openId string, count chan int) {
	stmt, err := mysqlApi.db.Prepare("insert into jmqjopenids (appid,openid,create_time) values (?,?,NOW())")
	if err != nil {
		log.Fatal(err)
	}
	v := <-count
	log.Printf("openid -> %s , count-> %d", openId, v)
	stmt.Exec(appId, openId)
	stmt.Close()
}

func (mysqlApi *MysqlApi) SaveOpenIds(appId string, openIds []string) {
	sql := "insert into jmqjopenids (appid,openid,create_time) values "
	for _, openId := range openIds {
		sql = sql + "('" + appId + "','" + openId + "',NOW()),"
	}
	sql = sql[0 : len(sql)-1]
	mysqlApi.db.Exec(sql)
}

func (mysqlApi *MysqlApi) SaveOpenIdsBySingle(appId string, openIds []string) {
	sql := "insert into jmqjopenids_test (appid,openid,create_time) values "
	for _, openId := range openIds {
		sql = sql + "('" + appId + "','" + openId + "',NOW()),"
	}
	sql = sql[0 : len(sql)-1]
	mysqlApi.db.Exec(sql)
}

func (mysqlApi *MysqlApi) DeleteByAppId(appId string) {
	stmt, err := mysqlApi.db.Prepare("delete from jmqjopenids where appid=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(appId)
	stmt.Close()
}


func (mysqlApi *MysqlApi) DeleteAppIdBySingle(appId string) {
	stmt, err := mysqlApi.db.Prepare("delete from jmqjopenids_test where appid=?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(appId)
	stmt.Close()
}
