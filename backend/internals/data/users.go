package data

import "database/sql"

type UserModel struct {
	db *sql.DB
}

func New(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}
