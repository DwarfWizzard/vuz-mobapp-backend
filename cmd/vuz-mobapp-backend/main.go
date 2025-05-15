package main

import (
	"context"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/logger"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"

	eduRepo "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/infrastructure/repository"
	eduUC "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/usecase"
	userHR "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/delivery/rest"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/auth"
	userRepo "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"
	userUC "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usecase"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	log := logger.Logger()
	log.Info("Start vuz mobapp backend service")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connectionUrl := os.Getenv("POSTGRES_CONNECTION_STRING")
	if len(connectionUrl) == 0 {
		log.Fatal("POSTGRES_CONNECTION_STRING can not be empty")
	}

	apiPort := os.Getenv("API_PORT")
	if _, err := strconv.Atoi(apiPort); err != nil {
		log.Fatal("API_PORT should be integer")
	}

	db, err := pggorm.NewDB(connectionUrl, log)
	if err != nil {
		log.Fatal("Connect to postgres error", zap.Error(err))
	}

	eduRepo := eduRepo.NewRepo(db)
	eduUC := eduUC.NewEducationUC(eduRepo, log)

	userRepo := userRepo.NewRepo(db)
	userUseCase := userUC.NewUserUC(userRepo, eduUC, log)

	authService := auth.NewJWTProvider(os.Getenv("JWT_SECRET"), userRepo, 15*time.Minute)

	userHandler := userHR.NewUserHandler(userUseCase, authService, log)

	router := echo.New()
	router.HTTPErrorHandler = response.HttpErrorHandler
	router.Server.BaseContext = func(l net.Listener) context.Context { return ctx }

	router.GET("/auth/login", userHandler.Login)

	apiGroup := router.Group("/v1")

	userGroup := apiGroup.Group("/user")
	userGroup.GET("/:id", userHandler.UserInfo)
}
