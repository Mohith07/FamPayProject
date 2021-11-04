package clients

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	g "gorm.io/gorm"

	"FamPayProject/config"
)

var (
	db *g.DB
)

func Init() {
	var err error
	log.Infof("connecting to the DB")
	conf := config.GetConfig()
	db, err = connectToDB(conf.DBUsername, conf.DBPassword,
		conf.DBHost, conf.DBPort,
		conf.DBName, conf.DBType,
		conf.DBSSLMode, true,
		10, 10)
	if err != nil {
		fmt.Println("error connecting to the DB " + err.Error())
		return
	}
}

func connectToDB(user, password, host, port, dbName, dbType, sslMode string, logMode bool, maxIdleConnections, maxOpenConnections int) (db *g.DB, err error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		user, password, host, port, dbName, sslMode)
	db, err = g.Open(postgres.Open(connectionString), &g.Config{})
	if err != nil {
		return nil, err
	}
	yDB, errNew := db.DB() // check the database connectivity
	if errNew != nil {
		return nil, errNew
	}
	yDB.SetMaxIdleConns(maxIdleConnections)
	yDB.SetMaxOpenConns(maxOpenConnections)
	return db, nil
}

// GetDB for getting HM database
func GetDB() *g.DB {
	return db
}