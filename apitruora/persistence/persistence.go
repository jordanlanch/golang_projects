package persistence

import (
	"errors"
	"log"
	"os"

	"apitruora/models/dbmodels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //for use postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   //for use sqlite at tests
)

var (
	ErrDBConnection = errors.New("failed to connect database")
)

//ConnectToDB func that create a DB connection
func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("TRACKING_DATABASE_URL"))
	return db, err
}

//CleanDB data base
func CleanDB() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(ErrDBConnection)
	}
	db.Exec("drop table server;")
	defer db.Close()
}

//MigrateDB database
func MigrateDB() (*gorm.DB, error) {
	db, err := ConnectToDB()

	if err != nil {
		utils.Error(err)
		panic(ErrDBConnection)
	}

	db.AutoMigrate(&dbmodels.Vehicles{})
	return db, err
}
