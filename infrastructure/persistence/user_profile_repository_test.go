package persistence_test

import (
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/achjailani/kickoff-go-rest/domain/entity"
	"github.com/achjailani/kickoff-go-rest/tests"
	"testing"
)

func TestUserProfileRepository_Save(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	userRepo := test.DBService.UserRepository
	userProfileRepo := test.DBService.UserProfileRepository

	f := faker.New()

	var (
		name = "Jay TEST001"
	)

	user := entity.User{
		Name:     name,
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(30),
	}

	t.Run("it should valid save user", func(t *testing.T) {
		err := userRepo.Save(ctx, &user)

		assert.NoError(t, err)
	})

	userProfile := entity.UserProfile{
		UserID: user.ID,
		Email:  f.Internet().Email(),
		Bio:    f.RandomLetter(),
	}

	t.Run("it should valid save", func(t *testing.T) {
		err := userProfileRepo.Save(ctx, &userProfile)

		assert.NoError(t, err)
	})

	userProfile2 := entity.UserProfile{
		Email: f.Internet().Email(),
		Bio:   f.RandomLetter(),
	}

	t.Run("it should error with empty user id", func(t *testing.T) {
		err := userProfileRepo.Save(ctx, &userProfile2)

		assert.Error(t, err)
	})
}

func TestUserProfileRepository_Update(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	userRepo := test.DBService.UserRepository
	userProfileRepo := test.DBService.UserProfileRepository

	f := faker.New()

	var (
		name = "Jay TEST001"
	)

	user := entity.User{
		Name:     name,
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(30),
	}

	t.Run("it should valid save user", func(t *testing.T) {
		err := userRepo.Save(ctx, &user)

		assert.NoError(t, err)
	})

	userProfile := entity.UserProfile{
		UserID: user.ID,
		Email:  f.Internet().Email(),
		Bio:    f.RandomLetter(),
	}

	t.Run("it should valid save", func(t *testing.T) {
		err := userProfileRepo.Save(ctx, &userProfile)

		assert.NoError(t, err)
	})

	t.Run("it should valid update", func(t *testing.T) {
		userId := user.ID
		userProfile = entity.UserProfile{
			Email: f.Internet().Email(),
		}

		err := userProfileRepo.Update(ctx, &userProfile, userId)

		assert.NoError(t, err)
	})
}

func TestUserProfileRepository_Delete(t *testing.T) {
	test := tests.Init()

	ctx := test.Ctx
	userRepo := test.DBService.UserRepository
	userProfileRepo := test.DBService.UserProfileRepository

	f := faker.New()

	var (
		name = "Jay TEST001"
	)

	user := entity.User{
		Name:     name,
		Username: f.RandomStringWithLength(10),
		Password: f.RandomStringWithLength(30),
	}

	t.Run("it should valid save user", func(t *testing.T) {
		err := userRepo.Save(ctx, &user)

		assert.NoError(t, err)
	})

	userProfile := entity.UserProfile{
		UserID: user.ID,
		Email:  f.Internet().Email(),
		Bio:    f.RandomLetter(),
	}

	t.Run("it should valid save", func(t *testing.T) {
		err := userProfileRepo.Save(ctx, &userProfile)

		assert.NoError(t, err)
	})

	t.Run("it should valid delete", func(t *testing.T) {
		userId := user.ID

		err := userProfileRepo.Delete(ctx, userId)

		assert.NoError(t, err)
	})
}
