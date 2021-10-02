package database

import (
	"fmt"
	"gits-echo-boilerplate/script/migration"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "gits-echo"
// )

func Init() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// sqlDB, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Hour)
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=gits-echo port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	migration.InitMigration(gormDB)

}

func CreateCon() *gorm.DB {
	return gormDB
}
