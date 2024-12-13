package db

import (
	"database/sql"
	_"github.com/lib/pq"
)

func Conexao() (*sql.DB, error){
	conn := "user=postgres password=arcesti123 dbname=users sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err!=nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}