package main

import (
	"context"
	"encoding/json"
	"strconv"
)

func initCache() {

	rdb.FlushAll(context.Background())

	var estateStructs []Estate
	var chairStructs []Chair

	if err := db.Select(&estateStructs, `SELECT * FROM estate`); err != nil {
		panic(err)
	}
	chairStructs = []Chair{}
	if err := db.Select(&chairStructs, `SELECT * FROM chair`); err != nil {
		panic(err)
	}
	for _, estate := range estateStructs {
		data, err := json.Marshal(&estate)
		if err != nil {
			panic(err)
		}
		status := rdb.Set(context.Background(), "estate_"+strconv.Itoa(int(estate.ID)), data, 0)
		if _, err := status.Result(); err != nil {
			panic(err)
		}
	}
	for _, chair := range chairStructs {
		data, err := json.Marshal(&chair)
		if err != nil {
			panic(err)
		}
		status := rdb.Set(context.Background(), "chair_"+strconv.Itoa(int(chair.ID)), data, 0)
		if _, err := status.Result(); err != nil {
			panic(err)
		}
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
