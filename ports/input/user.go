package input

import (
	"hexagonal-potato/application/repositories"
	"hexagonal-potato/domain"
	"hexagonal-potato/utils"
)

func CreateUser(user domain.AuthUser) string {
	if utils.ValidateCreateRequest(user) {
		err := repositories.CreateUserHandler(user)
		if err != nil {
			return "Success!"
		}
		return "Couldn't create user!"
	}
	return "Couldn't create user, Request is not valid"
}

func GetUser(user domain.AuthUser) domain.AuthUser {
	var returnedUser domain.AuthUser
	if utils.ValidateGetRequest(user) {
		err := repositories.GetUserHandler(user)
		if err != nil {
			return returnedUser
		}
		returnedUser = user
		return returnedUser
	} else {
		return returnedUser
	}

}

func UpdateUser(user domain.AuthUser) domain.AuthUser {
	var returnedUser domain.AuthUser
	if utils.ValidateUpdateRequest(user) {
		err := repositories.UpdateUserHandler(user)
		if err != nil {
			return GetUser(user)
		}
		return user
	}
	return returnedUser
}

func DeleteUser(user domain.AuthUser) bool {
	if utils.ValidateDeleteRequest(user) {
		err := repositories.DeleteUserHandler(user)
		return err != nil

	}
	return false

}

func ListUsers() map[string]domain.AuthUser {
	return repositories.ListUsers()
}
