package cmd

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/core/provider/connection"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	"github.com/achjailani/kickoff-go-rest/pkg/security"
	"github.com/achjailani/kickoff-go-rest/pkg/seeder"
	"log"
)

// NewCommand is a constructor
func NewCommand(conf *config.Config) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "create:secret",
			Usage: "A command to create secret key and public key",
			Action: func(c *cli.Context) error {
				secret, err := security.GenerateSecretKey()

				if err != nil {
					return err
				}

				fmt.Println("APP_PRIVATE_KEY:")
				fmt.Println(secret.PrivateKey)
				fmt.Println("APP_PUBLIC_KEY:")
				fmt.Println(secret.PublicKey)

				return nil
			},
		},
		{
			Name:  "db:migrate",
			Usage: "A command to migrate tables",
			Action: func(c *cli.Context) error {
				conn, err := connection.NewDBConnection(conf)
				if err != nil {
					return fmt.Errorf("unable to connect to db: %w", err)
				}

				registry := dao.NewRegistry()

				errMigrate := registry.AutoMigrate(conn)
				if errMigrate != nil {
					log.Fatalf("Unable to migrate tables, err: %v", errMigrate)
				}

				return nil
			},
		},
		{
			Name:  "db:seed",
			Usage: "A command to seed database",
			Action: func(c *cli.Context) error {
				conn, err := connection.NewDBConnection(conf)
				if err != nil {
					return fmt.Errorf("unable to connect to db: %w", err)
				}

				repo := dao.NewDBService(conn)

				ctx := context.Background()
				seed := seeder.NewSeeder(repo, &seeder.UserSeeder{})

				errSeed := seed.Run(ctx)
				if errSeed != nil {
					return fmt.Errorf("seed run: %w", errSeed)
				}

				return nil
			},
		},
	}
}
