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

	// 重複排除
	m := make(map[int64]struct{})
	for _, popularity := range popularities {
		m[popularity.Popularity] = struct{}{}
	}

	for popularity, _ := range m {
		if _, err := db.Exec(`UPDATE estate SET popularity_minus = ? WHERE popularity = ?`, popularity*-1, popularity); err != nil {
			panic(err)
		}
	}

}
