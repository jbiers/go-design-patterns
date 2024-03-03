package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type database struct {
}

var db *database

func GetDatabase() *database {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		// Double check for assuring goroutines are synchronized
		if db == nil {
			fmt.Println("Creating database instance.")
			db = &database{}
		} else {
			fmt.Println("Returning existing database instance.")
		}
	} else {
		fmt.Println("Returning existing database instance.")
	}

	return db
}
