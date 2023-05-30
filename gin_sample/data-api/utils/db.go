package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func connectDB(pgcon gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(pgcon, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// table name 為單數形式
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("Connect DB failed: ", err)
		panic(err)
	}
	// 如果沒有 schema 就自動建立，並設為 search_path, default: public
	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", "public"))
	db.Exec(fmt.Sprintf("SET search_path='%s'", "public"))

	return db.Session(&gorm.Session{PrepareStmt: true})
}

var DB *gorm.DB

func InitDB() {
	var pgcon = postgres.New(postgres.Config{
		DSN:                  CONFIG.DSN,
		PreferSimpleProtocol: true,
	})
	DB = connectDB(pgcon)
}

func DisconnectDB() {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
