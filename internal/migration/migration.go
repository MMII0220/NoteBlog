package migration

import (
	"fmt"
	// "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"myasd/config"
)

const (
	TableUser = `create table if not exists users (
		id serial primary key,
		full_name varchar(255) null,
		login varchar(255) not null,
		password varchar(255) not null,
		created_at timestamp default current_timestamp
	);`
	TableArticles = `create table if not exists articles (
		id serial primary key,
		name varchar(255) not null,
		content varchar(255) null,
		user_id int references users(id) not null,
		created_at timestamp default current_timestamp
	);`
)

func StartMigration() error {
	_, err := config.GetDBConnection().Exec(TableUser)
	if err != nil {
		return fmt.Errorf("error in creating table user %v: ", err)
	}
	_, err = config.GetDBConnection().Exec(TableArticles)
	if err != nil {
		return fmt.Errorf("error in creating table articles %v: ", err)
	}
	return nil
}
