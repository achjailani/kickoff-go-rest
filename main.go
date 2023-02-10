package main

import (
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/core/provider/connection"
	"github.com/achjailani/kickoff-go-rest/infrastructure/core/server"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/interface/cmd"
	"github.com/achjailani/kickoff-go-rest/interface/route"
	"github.com/achjailani/kickoff-go-rest/util"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var PORT = util.GetEnv("APP_PORT", "8181")

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("no .env file provided.")
	}

	conf := config.New()

	dbConn, errDBConn := connection.NewDBConnection(conf)
	if errDBConn != nil {
		log.Fatalf("Unable to connect database, err :%v", errDBConn)
	}

	redis := connection.NewRedisConnection(conf)

	jwt, errJWT := server.NewJWT(conf)
	if errJWT != nil {
		log.Fatalf("unable to initialize JWT, err: %v", errJWT)
	}

	databaseService := dao.NewDBService(dbConn)

	app := cmd.NewCli()
	app.Action = func(ctx *cli.Context) error {
		router := route.NewRouter(
			route.WithConfig(conf),
			route.WithRedis(redis),
			route.WithDatabaseService(databaseService),
			route.WithJWT(jwt),
		).Init()

		shutdownTimeout := 10 * time.Second

		err := server.RunHTTPServer(router, PORT, shutdownTimeout)
		if err != nil {
			return err
		}

		return nil
	}

	app.Commands = cmd.NewCommand(conf)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Unable to run CLI command, err: %v", err)
	}
}
