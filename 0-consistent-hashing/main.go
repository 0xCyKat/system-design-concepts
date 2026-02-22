package main

import (
	"consistent_hashing/database"
	"fmt"

	"github.com/kpechenenko/rword"
)

func main() {

	dbm := database.Init()

	for i := 0; i <= 4; i++ {
		db := database.Database{
			IP:   fmt.Sprintf("10.0.0.%d", i),
			Data: map[string]string{},
		}

		dbm.AddDatabase(db)
	}

	fmt.Println(dbm.GetSorted())

	gen, _ := rword.New()
	words := gen.WordList(50)

	for _, word := range words {
		dbm.InsertData(word, word+"_value")
		fmt.Println(dbm.GetData(word))
	}

	dbm.RemoveDatabase("10.0.0.4")

	fmt.Println("---------------- Removed Database ----------------")

	for _, word := range words {
		fmt.Println(dbm.GetData(word))
	}

	db0 := database.Database{
		IP:   "10.0.0.4",
		Data: map[string]string{},
	}

	dbm.AddDatabase(db0)

	fmt.Println("---------------- Reverted ----------------")

	for _, word := range words {
		fmt.Println(dbm.GetData(word))
	}
}
