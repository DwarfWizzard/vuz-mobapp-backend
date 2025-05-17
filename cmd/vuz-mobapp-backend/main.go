package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/logger"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"

	eduHR "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/delivery/rest"
	eduRepo "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/infrastructure/repository"
	eduUC "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/usecase"
	userHR "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/delivery/rest"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/auth"
	userRepo "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"
	userUC "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usecase"

	eduMigration "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/infrastructure/repository/migration"
	userMigration "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository/migration"

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

	db, err := pggorm.NewDB(connectionUrl)
	if err != nil {
		log.Fatal("Connect to postgres error", zap.Error(err))
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("Parse location error", zap.Error(err))
	}

	makeMigrateAndSeed(log, db)

	eduRepo := eduRepo.NewRepo(db)
	eduUC := eduUC.NewEducationUC(eduRepo, loc, log)

	userRepo := userRepo.NewRepo(db)
	userUseCase := userUC.NewUserUC(userRepo, eduUC, log)

	authService := auth.NewJWTProvider(os.Getenv("JWT_SECRET"), userRepo, 15*time.Minute)

	userHandler := userHR.NewUserHandler(userUseCase, authService, log)
	eduHandler := eduHR.NewEducationHandler(eduUC, log)

	router := echo.New()
	router.HTTPErrorHandler = response.HttpErrorHandler
	router.Server.BaseContext = func(l net.Listener) context.Context { return ctx }

	router.POST("/auth/login", userHandler.Login)
	router.POST("/auth/refresh", userHandler.Refresh)

	apiGroup := router.Group("/v1")

	// /v1/user/...
	userGroup := apiGroup.Group("/user", userHandler.AuthMiddleware())
	userGroup.GET("", userHandler.UserInfo)

	// /v1/edu/...
	eduGroup := apiGroup.Group("/edu")

	eventsGroup := eduGroup.Group("/events")
	eventsGroup.GET("", eduHandler.ActiveEvents)
	eventsGroup.GET("/:id", eduHandler.EventInfo)

	groupsGroup := eduGroup.Group("/groups", userHandler.AuthMiddleware())
	groupsGroup.GET("/:id/schedule", eduHandler.GroupSchedule)

	log.Info("Start API server")

	go func() {
		if err := router.Start(":" + apiPort); err != nil {
			log.Warn("Start server error", zap.Error(err))
		}
	}()

	// Waiting for OS signal
	chSig := make(chan os.Signal, 1)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	sig := <-chSig

	log.Info(fmt.Sprintf("OS signal received: %s", sig))

	log.Info("Closing API server")
	if err := router.Close(); err != nil {
		log.Warn("Closing API server error", zap.Error(err))
	}

	log.Info("Service stopped")

	log.Sync()
}

func makeMigrateAndSeed(log *zap.Logger, dbClient *pggorm.Db) {
	err := userMigration.Migrate(dbClient)
	if err != nil {
		log.Fatal("User migration error", zap.Error(err))
	}

	err = eduMigration.Migrate(dbClient)
	if err != nil {
		log.Fatal("Edu migration error", zap.Error(err))
	}

	err = userMigration.Seed(dbClient)
	if err != nil {
		log.Fatal("User seed error", zap.Error(err))
	}

	err = eduMigration.Seed(dbClient)
	if err != nil {
		log.Fatal("Edu seed error", zap.Error(err))
	}

	log.Info("Migrate and seed successfully made")
}
