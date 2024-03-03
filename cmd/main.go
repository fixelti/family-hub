package main

import (
	"context"
	"fmt"

	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres"
	httpTransport "github.com/fixelti/family-hub/internal/transport/http"
	"github.com/fixelti/family-hub/internal/usecase"
	libPostgres "github.com/fixelti/family-hub/lib/database/postgres"
	"github.com/fixelti/family-hub/lib/logger/zap"
	zapLib "go.uber.org/zap"
)

func main() {
	config := config.New("./")
	logger := zap.New(config.Debug)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db := libPostgres.New(context.Background(), dsn)
	repositoryManager := postgres.New(db, logger)
	usecaseManager := usecase.New(config, logger, repositoryManager.User)

	server := httpTransport.New(usecaseManager.User, config, logger)

	logger.Info(fmt.Sprintf("start server on port: %s\n", config.Server.Port))
	if err := server.Start(config.Server.Port); err != nil {
		logger.Error("failed to start server", zapLib.Error(err))
	}
}
