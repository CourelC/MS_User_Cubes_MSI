package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// DB params
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

var DB *gorm.DB

func GormOpen() (gormDB *gorm.DB, err error) {
	var infoDatabase infoDatabaseSQL
	err = infoDatabase.getDiverConn("Databases.MySQL.User_MS_Db_Cubes")
	if err != nil {
		return nil, err
	}
	gormDB, err = gorm.Open(mysql.Open(infoDatabase.Write.DriverConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return
	}

	err = gormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(infoDatabase.Read.DriverConn)},
	}))

	return
}
