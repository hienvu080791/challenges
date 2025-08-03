package storage

import (
	"demo_challenges/source/module/utils"
	"fmt"
	"sync"
)

var userStore = sync.Map{}

type User struct {
	Username string
	Password string
}

func CreateUser(user User) error {
	hashedPassword, err := utils.HashedPassword(user.Password)
	if err != nil {
		return err
	}
	_, isExist := userStore.LoadOrStore(user.Username, hashedPassword)
	if isExist {
		return fmt.Errorf("username %s already exists. Please try another one", user.Username)
	}
	//fmt.Println("Current users in UserStore:")
	//userStore.Range(func(key, value interface{}) bool {
	//	fmt.Printf("Username: %s, HashedPassword: %s\n", key, value)
	//	return true
	//})
	return nil
}

func CheckLogin(user User) error {
	storedPassword, ok := userStore.Load(user.Username)
	if !ok {
		return fmt.Errorf("user not exists")
	}
	if utils.ComparePassword(storedPassword, user.Password) != nil {
		return fmt.Errorf("incorrect password")
	}
	return nil
}
