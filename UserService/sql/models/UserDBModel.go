package models

import (
	"UserService/config"
	"UserService/errors"
	"UserService/sql/UserDB"
	_struct "UserService/struct"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	Config := config.Instance
	username := Config.GetConfig("username")
	password := Config.GetConfig("password")
	dbName := Config.GetConfig("dbName")
	protocol := Config.GetConfig("protocol")
	host := Config.GetConfig("MySqlHost")
	port := Config.GetConfig("MySqlPort")
	connection := fmt.Sprintf("%v:%v@%v(%v:%v)/%v", username, password, protocol, host, port, dbName)
	fmt.Printf("Connecting to database %v\n", connection)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatalf("Cannot connect to database %v\n", dbName)
		log.Fatal(err)
	}
	Queries := UserDB.New(db)
	Instance = &UserDBModel{
		Config:  Config,
		Db:      db,
		Queries: Queries,
	}
}

var Instance *UserDBModel

type UserDBModel struct {
	Config  config.Config
	Db      *sql.DB
	Queries *UserDB.Queries
}

func (u *UserDBModel) RegisterUser(context context.Context, request _struct.RegisterUserParams) (int32, int32, error) {
	tx, err := u.Db.Begin()
	if err != nil {
		log.Fatalf("%v\n", err)
		return errors.ERROR_WHEN_CREATE_TRANSACTION, -1, nil
	}
	defer tx.Rollback()
	qtx := u.Queries.WithTx(tx)
	res, err := qtx.CreateUser(context, request.Username)
	if err != nil {
		return errors.ERROR_WHEN_SAVE_USER, -1, nil
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return errors.ERROR_WHEN_SAVE_USER, -1, nil
	}
	UserId := sql.NullInt32{
		Int32: int32(lastInsertId),
		Valid: true,
	}
	err = qtx.AddPassword(context, UserDB.AddPasswordParams{
		Userid:   UserId,
		Password: request.Password,
	})
	if err != nil {
		return errors.ERROR_WHEN_SAVE_PASSWORD, -1, nil
	}
	err = tx.Commit()
	if err != nil {
		return errors.ERROR_WHEN_COMMIT_REGISTER_USER, -1, nil
	}
	return errors.SUCCESS, int32(lastInsertId), nil
}
