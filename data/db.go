package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	// user=xx, password=xx dbname=xx 切换为自己的postgres数据库用户、密码名和数据库名
	Db, err = sql.Open("postgres", "user=gwp password=741213 dbname=ChitChat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
