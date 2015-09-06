package db

import (
	glog "code.google.com/p/log4go"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// 数据库操作句柄.
var MySQL *sql.DB

// 初始化数据库连接.
func InitDB() {

	glog.Info("Connecting DB....")
	var err error
	MySQL, err = sql.Open("mysql", Conf.AppDBURL)

	if nil != err {
		glog.Error(err)
		os.Exit(-1)
	}

	// 实际测试一次
	test := 0
	if err := MySQL.QueryRow("SELECT 1").Scan(&test); err != nil {
		glog.Error(err)
		os.Exit(-1)
	}

	glog.Debug("DB max idle conns [%d]", Conf.AppDBMaxIdleConns)
	glog.Debug("DB max open conns [%d]", Conf.AppDBMaxOpenConns)

	MySQL.SetMaxIdleConns(Conf.AppDBMaxIdleConns)
	MySQL.SetMaxOpenConns(Conf.AppDBMaxOpenConns)

	glog.Debug("DB connected")
}

// 关闭数据库连接.
func CloseDB() {
	MySQL.Close()
}
