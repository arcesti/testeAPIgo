package db

import (
	_"fmt"

	"github.com/arcesti/PrimeiraAPIgo/user"
)

func Create(usr user.User)error {
	sql := "INSERT INTO users (nome, email) VALUES($1,$2)"
	db, err := Conexao();
	if err!=nil {
		return err
	}
	_,err = db.Exec(sql, usr.Nome, usr.Email)
	return err
}