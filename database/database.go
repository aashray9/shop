package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Connection() *gorm.DB {
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	dsn := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	dsn_M2 := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	dsn_M3 := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	dsn_M5 := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	dsn_M6 := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	dsn_M8 := DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{mysql.Open(dsn_M2)},
		Replicas: []gorm.Dialector{mysql.Open(dsn_M3)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}).Register(dbresolver.Config{
		// use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
		Replicas: []gorm.Dialector{mysql.Open(dsn_M5)},
	}, "lead_master").Register(dbresolver.Config{
		// use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
		Sources:  []gorm.Dialector{mysql.Open(dsn_M6)},
		Replicas: []gorm.Dialector{mysql.Open(dsn_M8)},
	}, "orders", "secondary"))
	DB = db
	return DB

}

func GetConnection() *gorm.DB {
	return DB
}
