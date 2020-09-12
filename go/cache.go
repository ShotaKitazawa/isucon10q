package main

import "sync"

var estateStructs []Estate
var estateMap map[int64]Estate
var estateLock sync.Mutex
var chairStructs []Chair
var chairMap map[int64]Chair
var chairLock sync.Mutex

func initCache() {
	estateMap = make(map[int64]Estate, 100000)
	chairMap = make(map[int64]Chair, 100000)
	estateStructs = []Estate{}
	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
	chairStructs = []Chair{}
	if err := db.Select(&chairStructs, `SELECT * FROM chair`); err != nil {
		panic(err)
	}
	for _, estate := range estateStructs {
		estateMap[estate.ID] = estate
	}
	for _, chair := range chairMap {
		chairMap[chair.ID] = chair
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
