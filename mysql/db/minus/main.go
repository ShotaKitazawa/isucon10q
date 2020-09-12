package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", "isucon", "isucon", "127.0.0.1", 3306, "isuumo")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	var popularities []struct {
		Popularity int64 `db:"popularity"`
	}
	db.Select(&popularities, `SELECT popularity FROM estate`)

	for _, popularity := range popularities {
		if _, err := db.Exec(`UPDATE estate SET popularity_minus = ? WHERE popularity = ?`, popularity.Popularity*-1, popularity.Popularity); err != nil {
			panic(err)
		}
	}

}
