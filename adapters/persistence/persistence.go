package persistence

import (
	"errors"
	"hexagonal-potato/domain"
	"sync"

	"github.com/google/uuid"
)

var (
	users      = make(map[string]domain.AuthUser)
	UserID     = uuid.NewString()
	UsersMutex = sync.Mutex{}
)

func CreateUser(user domain.AuthUser) error {
	UsersMutex.Lock()
	_, exists := users[user.ID]
	UsersMutex.Unlock()

	if exists {
		return errors.New("user already exists")
	}

	UsersMutex.Lock()
	users[user.ID] = user
	UsersMutex.Unlock()
	return nil
}

func GetUser(user domain.AuthUser) error {
	UsersMutex.Lock()
	_, exists := users[user.ID]
	UsersMutex.Unlock()

	if !exists {
		return errors.New("user doesn't exist")
	}
	return nil
}

func UpdateUser(user domain.AuthUser) error {
	UsersMutex.Lock()
	_, exists := users[user.ID]
	UsersMutex.Unlock()
	if !exists {
		return errors.New("user doesn't exist")
	}
	UsersMutex.Lock()
	users[user.ID] = user
	UsersMutex.Unlock()
	return nil
}

func DeleteUser(user domain.AuthUser) error {
	UsersMutex.Lock()
	_, exists := users[user.ID]
	UsersMutex.Unlock()
	if !exists {
		return errors.New("user doesn't exist")
	}
	UsersMutex.Lock()
	delete(users, user.ID)
	UsersMutex.Unlock()
	return nil
}

func ListUsers() map[string]domain.AuthUser {
	return users
}
