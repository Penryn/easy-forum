package service

import "easy-forum/internal/model"

func GetUserByUsername(username string) (user *model.User, err error) {
	user, err = d.GetUserByUsername(ctx, username)
	return user, err
}

func CreateUser(username string,name string, password string,utype int) (err error) {
	err = d.CreateUser(ctx, &model.User{
		Username: username,
		Name: name,
		Password: password,
		Type: utype,
	})
	return err
}