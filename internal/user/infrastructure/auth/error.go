package auth

import "errors"

var (
	errSigningMethod = errors.New("invalid signing method")
	errInvalidToken  = errors.New("invalid token")
	errTokenExpired  = errors.New("token expired")
)

func ErrorIsInvalidSigningMethod(err error) bool {
	return errors.Is(err, errSigningMethod)
}

func ErrorIsInvalidToken(err error) bool {
	return errors.Is(err, errInvalidToken)
}

func ErrorIsTokenExpired(err error) bool {
	return errors.Is(err, errTokenExpired)
}
