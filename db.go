package main


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB()(*sql.DB, error){

db, err := sql.Open("mysql", "cx70785_bitrix:Dsuheprf<24@tcp(46.138.170.253)/cx70785_bitrix")

	if err != nil{
		return nil, err
	}

	err = db.Ping()
		if err != nil{
			return nil, err
		}
		fmt.Println("Успешное подключение к БД")
		return db, nil
}
