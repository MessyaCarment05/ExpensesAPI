package database

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetConnection(){
	log.Println("Call GetConnection")
	var err error
	dsn :="root:@tcp(localhost:3306)/expenses_api?parseTime=true"
	DB,err=sql.Open("mysql", dsn)

	if err!=nil{
		log.Fatal(err)
	}

	// cek db bisa pakai atau gak pakai ping
	errPing:=DB.Ping()
	if errPing!=nil{
		log.Fatal(err)
	}

	DB.SetConnMaxIdleTime(10)
	DB.SetMaxOpenConns(50)
	DB.SetConnMaxIdleTime(5 *time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	log.Println("Connected to DB")
}