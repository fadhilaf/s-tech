package main

import (
	"github.com/fadhilaf/s-tech/common/env"

	"github.com/fadhilaf/s-tech/common/postgres"

	"github.com/fadhilaf/s-tech/internal/app"
)

func main() {
	appConfig := env.LoadConfig(".env")
	
	postgresDb := postgres.Start(appConfig.PostgresUrl, appConfig.PostgresMigratePath, appConfig.IsMigrateInit)

	app := app.New(appConfig, postgresDb)

	app.StartServer()
}
