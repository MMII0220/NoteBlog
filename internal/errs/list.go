package errs

import "errors"

var (
	ErrNotFound                = errors.New("not found")
	ErrArticleNotFound         = errors.New("article not found")
	ErrUserIDNotFoundInContext = errors.New("user id not found in context")

	ErrDuplicateArticle = errors.New("article already exists")
	ErrInvalidArticle   = errors.New("invalid data article")
	ErrUserNotExists    = errors.New("user does not exist")

	ErrUsernameAlreadyExists    = errors.New("username already exists")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")
	ErrFillRequiredFields       = errors.New("please fill in all required fields")

	ErrIncorrectRefreshToken = errors.New("incorrect refresh token")

	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidPathParam   = errors.New("invalid path param")
)
