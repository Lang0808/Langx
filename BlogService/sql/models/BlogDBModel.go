package models

import (
	"BlogService/config"
	"BlogService/errors"
	"BlogService/sql/BlogDB"
	_struct "BlogService/struct"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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
	Queries := BlogDB.New(db)
	Instance = &BlogDBModel{
		Config:  Config,
		Db:      db,
		Queries: Queries,
	}
}

var Instance *BlogDBModel

type BlogDBModel struct {
	Config  config.Config
	Db      *sql.DB
	Queries *BlogDB.Queries
}

func (u *BlogDBModel) CreatePost(context context.Context, request _struct.CreatePostParams) (int32, *_struct.CreatePostReturn, error) {
	res, err := u.Queries.CreateBlog(context, BlogDB.CreateBlogParams{
		Authorid: request.AuthorId,
		Title:    request.Content,
		Content:  request.Content,
		Active:   false,
		Timecreate: sql.NullInt64{
			Int64: time.Now().Unix(),
			Valid: true,
		},
	})
	if err != nil {
		return errors.ERROR_WHEN_SAVE_POST, nil, err
	}
	PostId, err := res.LastInsertId()
	if err != nil {
		return errors.ERROR_WHEN_GET_INSERTED_ID, nil, err
	}
	return errors.SUCCESS, &_struct.CreatePostReturn{
		PostID: int32(PostId),
	}, nil
}
