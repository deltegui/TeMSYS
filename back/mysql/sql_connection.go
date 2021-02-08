package mysql

import (
	"fmt"
	"log"
	"temsys/configuration"

	// Required to connect to msyql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Connect to msyql using a configuration and checks the connection.
// If it cant connect to mysql, it panics. There is no reason to
// continue if you dont have database.
func Connect(config configuration.Configuration) *sqlx.DB {
	db := connect(config)
	checkConnection(db, config)
	return db
}

func connect(config configuration.Configuration) *sqlx.DB {
	db, err := sqlx.Open(config.DatabaseDriver, fmt.Sprintf("%s?parseTime=true", config.Database))
	if err != nil {
		log.Fatalln("Error creating connection to database: ", err)
	}
	log.Printf("Connected to %s, using %s\n", config.Database, config.DatabaseDriver)
	return db
}

func checkConnection(conn *sqlx.DB, config configuration.Configuration) {
	if err := conn.Ping(); err != nil {
		log.Fatalln("Error checking connection to database: ", err)
	}
}
