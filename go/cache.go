package main

import (
	"context"
	"encoding/json"
	"strconv"
)

type ChairForCache struct {
	ID              int64  `db:"id" json:"id"`
	Name            string `db:"name" json:"name"`
	Description     string `db:"description" json:"description"`
	Thumbnail       string `db:"thumbnail" json:"thumbnail"`
	Price           int64  `db:"price" json:"price"`
	Height          int64  `db:"height" json:"height"`
	Width           int64  `db:"width" json:"width"`
	Depth           int64  `db:"depth" json:"depth"`
	Color           string `db:"color" json:"color"`
	Features        string `db:"features" json:"features"`
	Kind            string `db:"kind" json:"kind"`
	Popularity      int64  `db:"popularity" json:"popularity"`
	Stock           int64  `db:"stock" json:"stock"`
	PopularityMinus int64  `db:"popularity_minus" json:"-"`
}

type EstateForCache struct {
	ID              int64   `db:"id" json:"id"`
	Thumbnail       string  `db:"thumbnail" json:"thumbnail"`
	Name            string  `db:"name" json:"name"`
	Description     string  `db:"description" json:"description"`
	Latitude        float64 `db:"latitude" json:"latitude"`
	Longitude       float64 `db:"longitude" json:"longitude"`
	Address         string  `db:"address" json:"address"`
	Rent            int64   `db:"rent" json:"rent"`
	DoorHeight      int64   `db:"door_height" json:"doorHeight"`
	DoorWidth       int64   `db:"door_width" json:"doorWidth"`
	Features        string  `db:"features" json:"features"`
	Popularity      int64   `db:"popularity" json:"popularity"`
	PopularityMinus int64   `db:"popularity_minus" json:"-"`
}

func initCache() {

	rdb.FlushAll(context.Background())

	var estateStructs []EstateForCache
	var chairStructs []ChairForCache

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
