package dbclient

import "database/sql"

var dbClient *sql.DB

func Init() {
	var err error
	dbClient, err = sql.Open("mysql", "root:theavengers2@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		panic(err.Error())
	}
}

func GetDBClient() *sql.DB {
	return dbClient
}
