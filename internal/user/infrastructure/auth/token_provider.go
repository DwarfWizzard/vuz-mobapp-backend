package auth

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"

	"github.com/golang-jwt/jwt"
)

type UserRepo interface {
	GetUserById(userId uint32) (*model.User, error)
}

type TokenProvider interface {
	GetApikeyByToken(ctx context.Context, tokenValue string) (*Apikey, error)
	GenerateUserTokenPair(ctx context.Context, user *model.User) (*TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error)
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
	Type   string
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
		vErr, _ := err.(*jwt.ValidationError)
		if vErr.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, errTokenExpired
		}

		return nil, vErr.Inner
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errInvalidToken
	}

	if claims.Type != "access" {
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

// TODO: add different token secrets
func (tp *jwtTokenProvider) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errSigningMethod
		}
		return []byte(tp.secret), nil
	})
	if err != nil {
		// Обработка ошибок просроченного токена и других проблем
		vErr, _ := err.(*jwt.ValidationError)
		if vErr != nil && vErr.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, errTokenExpired
		}
		return nil, errInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errInvalidToken
	}

	if claims.Type != "refresh" {
		return nil, errInvalidToken
	}

	user, err := tp.repo.GetUserById(claims.UserId)
	if err != nil {
		if repository.ErrorIsNoRows(err) {
			return nil, errInvalidToken
		}
		return nil, err
	}

	// Генерация новой пары токенов
	return tp.GenerateUserTokenPair(ctx, user)
}

func (tp *jwtTokenProvider) GenerateUserTokenPair(ctx context.Context, user *model.User) (*TokenPair, error) {
	now := time.Now().UTC()
	expiration := now.Add(tp.ttl)

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			IssuedAt:  now.Unix(),
		},
		"access",
		user.ID,
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		jwt.StandardClaims{
			ExpiresAt: expiration.Add(5 * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
		},
		"refresh",
		user.ID,
	})

	var signErr error

	accessToken, signErr := access.SignedString([]byte(tp.secret))
	if signErr != nil {
		return nil, signErr
	}

	refreshToken, signErr := refresh.SignedString([]byte(tp.secret))
	if signErr != nil {
		return nil, signErr
	}

	return &TokenPair{
		Token:   accessToken,
		Refresh: refreshToken,
	}, nil
}
