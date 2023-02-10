package exception

import "errors"

var (
	// ErrorBadRequest represents error bad request
	ErrorBadRequest = errors.New("common.bad_request")

	// ErrorUnprocessableEntity represents error unprocessable entity
	ErrorUnprocessableEntity = errors.New("common.unprocessable_entity")

	// ErrorInternalServerError represents error internal server
	ErrorInternalServerError = errors.New("common.internal_server_error")

	// ErrorForbidden represents error unauthorized
	ErrorForbidden = errors.New("common.forbidden")

	// ErrorUnauthorized represents error unauthorized
	ErrorUnauthorized = errors.New("common.unauthorized")

	// ErrInvalidToken represents error of invalid token
	ErrInvalidToken = errors.New("common.invalid_token")

	// ErrExpiredToken represents error of expired token
	ErrExpiredToken = errors.New("common.expired_token")
)
