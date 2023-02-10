package tests

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/infrastructure/core/provider/connection"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
	seeder2 "github.com/achjailani/kickoff-go-rest/pkg/seeder"
	"github.com/achjailani/kickoff-go-rest/tests/database"
	"github.com/achjailani/kickoff-go-rest/util"
	"log"
	"os"
)

// TestSuite is a struct represents its self
type TestSuite struct {
	Config    *config.Config
	DBService *dao.Repositories
	Ctx       context.Context
}

// Init is a constructor of test suite
func Init() *TestSuite {
	errT := os.Setenv("TEST_MODE", "true")
	if errT != nil {
		log.Fatalf("unable to set test mode, err: %w", errT)
	}

	if err := godotenv.Load(fmt.Sprintf("%s/.env", util.RootDir())); err != nil {
		log.Fatalf("no .env file provided.")
	}

	conf := config.New()

	dbConn, errDBConn := connection.NewDBConnection(conf)
	if errDBConn != nil {
		log.Fatalf("unable to connect to database, %v", errDBConn)
	}

	repo := dao.NewDBService(dbConn)
	ctx := context.Background()

	drop := database.NewDrop(dbConn)

	errDBReset := drop.Reset(ctx)
	if errDBReset != nil {
		log.Fatalf("unable to reset database, %v", errDBReset)
	}

	seed := seeder2.NewSeeder(repo, &seeder2.UserSeeder{})
	errSeed := seed.Run(ctx)

	if errSeed != nil {
		log.Fatalf("unable to run seeder %w.", errSeed)
	}

	return &TestSuite{
		Config:    conf,
		DBService: repo,
		Ctx:       ctx,
	}
}

// IsTestMode is function to check where execution on test mode (go test)
func IsTestMode() bool {
	return flag.Lookup("test.v").Value.(flag.Getter).Get().(bool)
}
