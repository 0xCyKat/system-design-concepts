package database

import (
	"crypto/sha256"
	"encoding/binary"
	"hash"
	"sort"
)

type DBManager struct {
	databases map[int]*Database
	hasher    hash.Hash
}

func Init() *DBManager {
	dbm := &DBManager{
		databases: map[int]*Database{},
		hasher:    sha256.New(),
	}

	return dbm
}

func (dbm *DBManager) getIndex(key string) int {
	dbm.hasher.Reset()
	dbm.hasher.Write([]byte(key))
	hashBytes := dbm.hasher.Sum(nil)

	return int(binary.BigEndian.Uint64(hashBytes[:8]))

}

func (dbm *DBManager) getSortedKeys() []int {
	keys := make([]int, 0)
	for k := range dbm.databases {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	return keys
}

func (dbm *DBManager) InsertData(key, value string) string {
	index := dbm.getIndex(key)
	keys := dbm.getSortedKeys()

	for _, val := range keys {
		if index < val {
			dbm.databases[val].PutData(key, value)
			return dbm.databases[val].IP
		}
	}

	dbm.databases[keys[0]].PutData(key, value)

	return dbm.databases[keys[0]].IP
}

func (dbm *DBManager) GetData(key string) (string, string) {
	index := dbm.getIndex(key)
	keys := dbm.getSortedKeys()

	for _, val := range keys {
		if index < val {
			return dbm.databases[val].GetData(key), dbm.databases[val].IP
		}
	}

	return dbm.databases[keys[0]].GetData(key), dbm.databases[keys[0]].IP
}

func (dbm *DBManager) AddDatabase(db Database) {
	index := dbm.getIndex(db.IP)
	dbm.databases[index] = &db

	keys := dbm.getSortedKeys()
	n := len(keys)

	if n == 1 {
		return
	}

	var nextInd int
	var prevInd int

	for i, key := range keys {
		if key == index {
			nextInd = keys[(i+1)%n]
			prevInd = keys[(i-1+n)%n]
			break
		}
	}

	nextDB := dbm.databases[nextInd]

	if len(nextDB.Data) == 0 {
		return
	}

	for k, v := range nextDB.Data {
		kInd := dbm.getIndex(k)

		var belongs bool
		if prevInd < index {
			belongs = kInd > prevInd && kInd <= index
		} else {
			belongs = kInd > prevInd || kInd <= index
		}

		/*
			Explanation:
			[10.0.0.3 10.0.0.0 10.0.0.2 10.0.0.1 10.0.0.4] - Sorted order based on Hash of IP (Circular array)

			Scenario 1: 10.0.0.4 is removed & added back
				- After removal - all the keys in (10.0.0.1, 10.0.0.3] should be in 10.0.0.3
				- Added back - all the keys in (10.0.0.1, 10.0.0.4] should be routed back to 10.0.0.4
					- For this to happen, check the nextDB and find all the keys, which lie in between them
						- kInd > prevInd && kInd <= index
						- Since here prevInd is 10.0.0.1 (index 3)
		*/

		if belongs {
			dbm.databases[index].PutData(k, v)
			delete(nextDB.Data, k)
		}
	}

}

func (dbm *DBManager) RemoveDatabase(IP string) {
	var db *Database

	for _, v := range dbm.databases {
		if v.IP == IP {
			db = v
			break
		}
	}

	if db == nil {
		return
	}

	index := dbm.getIndex(db.IP)
	keys := dbm.getSortedKeys()
	n := len(keys)

	if n == 1 {
		delete(dbm.databases, index)
		return
	}

	var nextInd int

	for i, key := range keys {
		if index == key {
			nextInd = keys[(i+1)%n]
			break
		}
	}

	nextDB := dbm.databases[nextInd]

	for k, v := range db.Data {
		nextDB.PutData(k, v)
		delete(db.Data, k)
	}

	delete(dbm.databases, index)

}

func (dbm *DBManager) GetSorted() []string {
	var sortedIPs []string

	keys := dbm.getSortedKeys()

	for _, key := range keys {
		ip := dbm.databases[key].IP

		sortedIPs = append(sortedIPs, ip)
	}

	return sortedIPs
}
