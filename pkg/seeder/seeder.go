package seeder

import (
	"context"
	"fmt"
	"github.com/achjailani/kickoff-go-rest/infrastructure/dao"
)

// Seeder is an interface should be implemented
type Seeder interface {
	Seed(ctx context.Context, repo *dao.Repositories) error
}

// AutoSeeder is a struct
type AutoSeeder struct {
	repo  *dao.Repositories
	seeds []Seeder
}

// NewSeeder is a constructor
func NewSeeder(repo *dao.Repositories, seeds ...Seeder) *AutoSeeder {
	registry := &AutoSeeder{repo: repo}

	for _, seeder := range seeds {
		registry.register(seeder)
	}

	return registry
}

// register is method to register new seeder
func (s *AutoSeeder) register(seeder Seeder) {
	s.seeds = append(s.seeds, seeder)
}

// Run is method to register all registered seeders
func (s *AutoSeeder) Run(ctx context.Context) error {
	for _, seeder := range s.seeds {
		err := seeder.Seed(ctx, s.repo)
		if err != nil {
			return fmt.Errorf("seeder seed: %w", err)
		}
	}

	return nil
}
