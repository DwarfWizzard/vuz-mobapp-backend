package auth

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usercore"
	"github.com/golang-jwt/jwt"
)

type UserRepo interface {
	GetUserById(userId uint32) (*usercore.User, error)
}

type TokenProvider interface {
	GetApikeyByToken(ctx context.Context, tokenValue string) (*Apikey, error)
	GenerateUserTokenPair(ctx context.Context, user *usercore.User) (*TokenPair, error)
}

type jwtTokenProvider struct {
	repo   UserRepo
	secret string
	ttl    time.Duration
}

func NewJWTProvider(secret string, repo UserRepo, ttl time.Duration) TokenProvider {
	return &jwtTokenProvider{secret: secret, repo: repo, ttl: ttl}
}

type Claims struct {
	jwt.StandardClaims
	UserId uint32
}

func (tp *jwtTokenProvider) GetApikeyByToken(ctx context.Context, tokenValue string) (*Apikey, error) {
	token, err := jwt.ParseWithClaims(tokenValue, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errSigningMethod
		}

		return []byte(tp.secret), nil
	})

	if err != nil {
		vErr, _ := err.(jwt.ValidationError)
		return nil, vErr.Inner
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errInvalidToken
	}

	user, err := tp.repo.GetUserById(claims.UserId)
	if err != nil {
		if repository.ErrorIsNoRows(err) {
			return nil, errInvalidToken
		}

		return nil, err
	}

	apikey := &Apikey{
		UserId: user.ID,
	}

	return apikey, nil
}

func (tp *jwtTokenProvider) GenerateUserTokenPair(ctx context.Context, user *usercore.User) (*TokenPair, error) {
	now := time.Now().UTC()
	expiration := now.Add(tp.ttl)

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			IssuedAt:  now.Unix(),
		},
		user.ID,
	})

	jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			IssuedAt:  now.Unix(),
		},
		user.ID,
	})

	var signErr error

	accessToken, signErr := access.SignedString([]byte(tp.secret))
	if signErr != nil {
		return nil, signErr
	}

	refreshToken, signErr := access.SignedString([]byte(tp.secret))
	if signErr != nil {
		return nil, signErr
	}

	return &TokenPair{
		Token:   accessToken,
		Refresh: refreshToken,
	}, nil
}
