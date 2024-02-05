package main

import (
	"context"
	"fmt"

	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres"
	httpTransport "github.com/fixelti/family-hub/internal/transport/http"
	"github.com/fixelti/family-hub/internal/usecase"
	libPostgres "github.com/fixelti/family-hub/lib/database/postgres"
)

func main() {
	config := config.New("./config", "./.env")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db := libPostgres.New(context.Background(), dsn)
	repositoryManager := postgres.New(db)
	usecaseManager := usecase.New(config, repositoryManager.User)

	httpTransport.New(config.Server.Port, usecaseManager.User)
}
