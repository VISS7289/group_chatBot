package Mysql

import (
	"fmt"
	"go-chatbot/Settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(sql Settings.MySqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		sql.User,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.DBName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(sql.MaxOpenConns)
	db.SetMaxIdleConns(sql.MaxIdleConns)
	return nil
}

func Close() {
	_ = db.Close()
}
