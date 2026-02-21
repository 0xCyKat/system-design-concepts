package database

import (
	"encoding/binary"
	"hash"
	"sort"
)

type DBManager struct {
	Databases map[int]*Database
	Hasher    hash.Hash
}

func (dbm *DBManager) getIndex(key string) int {
	dbm.Hasher.Reset()
	dbm.Hasher.Write([]byte(key))
	hashBytes := dbm.Hasher.Sum(nil)

	return int(binary.BigEndian.Uint64(hashBytes[:8]))

}

func (dbm *DBManager) getSortedKeys() []int {
	keys := make([]int, len(dbm.Databases))
	for k := range dbm.Databases {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func (dbm *DBManager) InsertData(key, value string) {
	index := dbm.getIndex(key)
	keys := dbm.getSortedKeys()

	for _, val := range keys {
		if index < val {
			dbm.Databases[val].PutData(key, value)
			return
		}
	}

	dbm.Databases[keys[0]].PutData(key, value)
}

func (dbm *DBManager) GetData(key string) string {
	index := dbm.getIndex(key)
	keys := dbm.getSortedKeys()

	for _, val := range keys {
		if index < val {
			return dbm.Databases[val].GetData(key)
		}
	}

	return dbm.Databases[keys[0]].GetData(key)
}

func (dbm *DBManager) AddDatabase(db Database) {
	index := dbm.getIndex(db.IP)
	dbm.Databases[index] = &db
}

/*
	TODO: Implement data movement

	Scenario:
		1 	2	4 - Databases

	1. All keys in (2, 4] - will be in 4
	2. Add database 3 -> 	1	2	3	4
		- keys in (3, 4] - will stay in 4
		- keys in (2, 3] - need to be moved to 3 from 4

	Sol: Iterate over keys in 4 -> if a key hash is less than 3, move it to 3

	- Similar for Delete Database as well
*/
