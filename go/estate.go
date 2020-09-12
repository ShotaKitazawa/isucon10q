package main

import "github.com/k0kubun/pp"

var estateStructs []Estate

func initEstate() {
	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
	pp.Print(estateStructs)
	// estateStructs = []Estate{
	//	Estate{
	//	},
	// }
}
