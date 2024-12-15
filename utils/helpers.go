package utils

import (
	"hexagonal-potato/domain"
)

func ValidateCreateRequest(user domain.AuthUser) bool {
	return user.Username != "" && user.Password != ""
}

func ValidateGetRequest(user domain.AuthUser) bool {
	return user.ID != ""
}

func ValidateUpdateRequest(user domain.AuthUser) bool {
	return user.Username != "" && user.Password != ""
}

func ValidateDeleteRequest(user domain.AuthUser) bool {
	return user.ID != ""
}
