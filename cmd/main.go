package main

import (
	"fmt"

	"github.com/Avtoelon/api"
	"github.com/Avtoelon/config"
	"github.com/Avtoelon/pkg/logger"
	gormadapter "github.com/casbin/gorm-adapter/v2"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	_, err := gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("new adapter error", logger.Error(err))

	}
	server := api.New(api.Option{
		Conf:   cfg,
		Logger: log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
