package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/Ras96/clean-architecture-sample/2_interface/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type sqlConfig struct {
	user string
	pass string
	host string
	port string
	db   string
}

type SQLHandler struct {
	conn *gorm.DB
}

func NewSQLHandler() database.SQLHandler {
	var conf sqlConfig
	if conf.user = os.Getenv("DB_USER"); conf.user == "" {
		conf.user = "root"
	}
	if conf.pass = os.Getenv("DB_PASS"); conf.pass == "" {
		conf.pass = "pass"
	}
	if conf.host = os.Getenv("DB_HOST"); conf.host == "" {
		conf.host = "localhost"
	}
	if conf.port = os.Getenv("DB_PORT"); conf.port == "" {
		conf.port = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.user, conf.pass, conf.host, conf.port, conf.db)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &SQLHandler{conn}
}

// SQLHandler(構造体)がSQLHandler(インターフェース)を満たすためにメソッドを定義する
func (hl *SQLHandler) Find(out interface{}, where ...interface{}) database.SQLHandler {
	db := hl.conn.Find(out, where...)

	return &SQLHandler{conn: db}
}

func (hl *SQLHandler) First(out interface{}, where ...interface{}) database.SQLHandler {
	db := hl.conn.First(out, where...)

	return &SQLHandler{conn: db}
}

func (hl *SQLHandler) Create(value interface{}) database.SQLHandler {
	db := hl.conn.Create(value)

	return &SQLHandler{conn: db}
}

func (hl *SQLHandler) Error() error {
	return hl.conn.Error
}
