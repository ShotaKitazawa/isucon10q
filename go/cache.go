package main

import "fmt"

var estateStructs []Estate
var chairStructs []Chair

func initEstate() {
	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
	if err := db.Select(&chairStructs, `SELECT * FROM chair`); err != nil {
		panic(err)
	}
	fmt.Println(len(estateStructs))
	fmt.Println(len(chairStructs))
}
