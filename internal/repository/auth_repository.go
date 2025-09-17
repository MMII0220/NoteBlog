package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"myasd/config"
	// "myasd/internal/controller"
	"myasd/internal/models"
)

func CreateUser(user models.User) error {
	query := `insert into users (full_name, login, password) values ($1, $2, $3)`
	_, err := config.GetDBConnection().Exec(query, user.FullName, user.Login, user.Password)
	if err != nil {
		return fmt.Errorf("error inserting user %v: ", err)
	}
	return nil
}

func GetUserByLogin(login string) (models.User, error) {
	var user models.User
	query := `select id, full_name, login, password, created_at from users where login=$1`
	err := config.GetDBConnection().Get(&user, query, login)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}
