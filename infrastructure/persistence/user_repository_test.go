package persistence_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/tests"
	"strings"
	"testing"
)

func TestUserRepository_Save(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	repo := test.DBService
	userRepo := test.DBService.UserRepository

	f := faker.New()

	role := entity.Role{
		Name:        f.RandomStringWithLength(10),
		Code:        strings.ToUpper(f.RandomStringWithLength(10)),
		Description: f.Lorem().Word(),
	}

	_ = repo.RoleRepository.Save(ctx, &role)

	user := entity.User{
		RoleID:   role.ID,
		Name:     f.Person().Name(),
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(20),
	}

	t.Run("if valid save", func(t *testing.T) {
		err := userRepo.Save(ctx, &user)

		assert.NoError(t, err)
	})
}

func TestUserRepository_FindByID(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	repo := test.DBService

	f := faker.New()

	var (
		id = f.UUID().V4()
	)

	t.Run("if record not found", func(t *testing.T) {
		result, err := repo.UserRepository.FindByID(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, result)
	})

	role := entity.Role{
		Name:        f.RandomStringWithLength(10),
		Code:        strings.ToUpper(f.RandomStringWithLength(10)),
		Description: f.Lorem().Word(),
	}

	_ = repo.RoleRepository.Save(ctx, &role)

	user := entity.User{
		ID:       id,
		RoleID:   role.ID,
		Name:     f.Person().Name(),
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(20),
	}

	t.Run("if valid save", func(t *testing.T) {
		err := repo.UserRepository.Save(ctx, &user)

		assert.NoError(t, err)
	})

	t.Run("if valid find by id", func(t *testing.T) {
		result, err := repo.UserRepository.FindByID(ctx, id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.Role)
	})
}

func TestUserRepository_Update(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	userRepo := test.DBService.UserRepository

	f := faker.New()

	user := entity.User{
		Name:     f.Person().Name(),
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(20),
	}

	t.Run("if valid save", func(t *testing.T) {
		err := userRepo.Save(ctx, &user)

		assert.NoError(t, err)
	})

	t.Run("if valid update", func(t *testing.T) {
		userId := user.ID

		err := userRepo.Update(ctx, &entity.User{
			Name: f.Person().Name(),
		}, userId)

		assert.NoError(t, err)
	})
}
