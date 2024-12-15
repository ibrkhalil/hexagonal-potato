package repositories

import (
	"hexagonal-potato/adapters/persistence"
	"hexagonal-potato/domain"
)

func ListUsers() map[string]domain.AuthUser {
	return persistence.ListUsers()
}

func CreateUserHandler(user domain.AuthUser) error {
	return persistence.CreateUser(user)
}

func GetUserHandler(user domain.AuthUser) error {
	return persistence.GetUser(user)
}

func UpdateUserHandler(user domain.AuthUser) error {
	return persistence.UpdateUser(user)

}

func DeleteUserHandler(user domain.AuthUser) error {
	return persistence.DeleteUser(user)
}
