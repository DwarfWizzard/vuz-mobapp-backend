package auth

import "errors"

var (
	errSigningMethod = errors.New("invalid signing method")
	errInvalidToken  = errors.New("invalid token")
)

func ErrorIsInvalidSigningMethod(err error) bool {
	return errors.Is(err, errSigningMethod)
}

func ErrorIsInvalidToken(err error) bool {
	return errors.Is(err, errInvalidToken)
}
