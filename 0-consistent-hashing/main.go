package main

import (
	"consistent_hashing/database"
	"crypto/sha256"
	"fmt"
)

func main() {
	dbm := &database.DBManager{
		Databases: map[int]*database.Database{},
		Hasher:    sha256.New(),
	}

	db1 := database.Database{
		IP:   "10.0.0.1",
		Data: map[string]string{},
	}

	db2 := database.Database{
		IP:   "10.0.0.2",
		Data: map[string]string{},
	}

	dbm.AddDatabase(db1)
	dbm.AddDatabase(db2)

	dbm.InsertData("name", "sai_srinivas")
	fmt.Println(dbm.GetData("name"))

	dbm.InsertData("age", "22")
	fmt.Println(dbm.GetData("age"))
}
