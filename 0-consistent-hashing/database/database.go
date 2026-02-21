package database

type Database struct {
	IP   string
	Data map[string]string
}

func (db *Database) GetData(key string) string {
	return db.Data[key]
}

func (db *Database) PutData(key string, value string) {
	db.Data[key] = value
}
