package util

import "net/http"

const OK = 200

var (
	DBError                 = NewError(http.StatusInternalServerError, "Failed to use DB")
	HashError               = NewError(http.StatusInternalServerError, "Failed to hash password")
	BindError               = NewError(http.StatusBadRequest, "Failed to bind")
	CreateError             = NewError(http.StatusInternalServerError, "Failed to Create")
	UsernameOrPasswordError = NewError(http.StatusUnauthorized, "Invalid username or password")
	GenerateTokenError      = NewError(http.StatusInternalServerError, "Failed to generate token")
	TokenError              = NewError(http.StatusUnauthorized, "Invalid token")
)
