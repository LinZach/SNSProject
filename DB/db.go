package DB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const(
	userName = "root"
	password = "1q2w3e4r"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "jdbc:mysql://localhost"
)

var DB *sql.DB

func InitDB()  {

	var err error
	DB ,err = sql.Open("mysql", "root:1q2w3e4r@tcp(localhost:3306)/skycat")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)

	if err = DB.Ping(); err != nil {
		fmt.Println("db connect err")
		panic(err)
	}

	fmt.Println("db connect success")
}

func Close() {
	DB.Close()
}