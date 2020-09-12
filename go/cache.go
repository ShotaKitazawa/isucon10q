package main

import (
	"encoding/json"
	"strconv"
)

//var estateMap map[int64]Estate
//var estateLock sync.Mutex
//var chairMap map[int64]Chair
//var chairLock sync.Mutex

func initCache() {
	// estateMap = make(map[int64]Estate, 100000)
	// chairMap = make(map[int64]Chair, 100000)

	var estateStructs []Estate
	var chairStructs []Chair

	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
	if err := db.Select(&chairStructs, `SELECT * FROM chair`); err != nil {
		panic(err)
	}

	for _, estate := range estateStructs {
		data, err := json.Marshal(&estate)
		if err != nil {
			panic(err)
		}
		c := pool.Get()
		if _, err := c.Do("SET", "estate_"+strconv.Itoa(int(estate.ID)), data); err != nil {
			panic(err)
		}
		c.Close()
	}
	for _, chair := range chairStructs {
		data, err := json.Marshal(&chair)
		if err != nil {
			panic(err)
		}
		c := pool.Get()
		if _, err := c.Do("SET", "chair_"+strconv.Itoa(int(chair.ID)), data); err != nil {
			panic(err)
		}
		c.Close()
	}
}
