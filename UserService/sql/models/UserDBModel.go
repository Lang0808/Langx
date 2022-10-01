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
	params := UserDB.CreateUserParams{
		Username: request.Username,
		Isadmin:  request.IsAdmin,
	}
	res, err := qtx.CreateUser(context, params)
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

func (u *UserDBModel) LoginUser(ctx context.Context, request _struct.LoginUserParams) (int32, *_struct.LoginUserReturn) {
	ans := &_struct.LoginUserReturn{
		UserId:   -1,
		Username: "",
	}
	user, err := u.Queries.GetUser(ctx, request.Username)
	if err != nil {
		return errors.ERROR_WHEN_GET_USER, ans
	}
	password, err := u.Queries.GetPassword(ctx, user.ID)
	if err != nil {
		return errors.ERROR_WHEN_GET_PASSWORD, ans
	}
	if equals(password, request.Password) {
		return errors.SUCCESS, &_struct.LoginUserReturn{
			UserId:   user.ID.Int32,
			Username: user.Username,
			IsAdmin:  ans.IsAdmin,
		}
	}
	return errors.PASSWORD_INCORRECT, ans
}

func equals(password UserDB.Password, password2 string) bool {
	return password.Password == password2
}
