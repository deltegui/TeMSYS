package persistence

import (
	"fmt"
	"log"
	"sensorapi/src/configuration"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SqlxConnection struct {
	db     *sqlx.DB
	config configuration.Configuration
}

func NewSqlxConnection(config configuration.Configuration) SqlxConnection {
	return SqlxConnection{config: config}
}

func (sqlxConn *SqlxConnection) GetConnection() *sqlx.DB {
	if sqlxConn.db == nil {
		sqlxConn.connect()
		sqlxConn.checkConnection()
	}
	return sqlxConn.db
}

func (conn *SqlxConnection) connect() {
	config := conn.config
	db, err := sqlx.Open(config.DatabaseDriver, fmt.Sprintf("%s?parseTime=true", config.Database))
	if err != nil {
		log.Fatalln("Error creating connection to database: ", err)
	}
	log.Printf("Connected to %s, using %s\n", config.Database, config.DatabaseDriver)
	conn.db = db
}

func (conn SqlxConnection) checkConnection() {
	if err := conn.db.Ping(); err != nil {
		log.Fatalln("Error checking connection to database: ", err)
	}
}
