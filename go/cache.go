package main

import "sync"

var estateMap map[int64]Estate
var estateLock sync.Mutex
var chairMap map[int64]Chair
var chairLock sync.Mutex

func initCache() {
	estateMap = make(map[int64]Estate)
	chairMap = make(map[int64]Chair)

	var estateStructs []Estate
	var chairStructs []Chair

	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
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
