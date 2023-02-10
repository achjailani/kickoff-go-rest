package jwt_test

import (
	"encoding/base64"
	"fmt"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
	"github.com/achjailani/kickoff-go-rest/tests"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	test := tests.Init()

	privateKey, errPri := base64.StdEncoding.DecodeString(test.Config.AppPrivateKey)
	publicKey, errPub := base64.StdEncoding.DecodeString(test.Config.AppPublicKey)

	assert.NoError(t, errPri)
	assert.NoError(t, errPub)

	jwtInstance := jwt.NewJWT(privateKey, publicKey)
	payload := jwt.JWTPayload{
		Identifier: uuid.New().String(),
		Name:       "Test User",
	}

	now := time.Now()
	ttl := time.Minute * 5

	fmt.Println(now)
	fmt.Println(now.Add(time.Minute * 5))

	jwtToken, err := jwtInstance.Create(ttl, &payload)

	t.Run("it should valid create token", func(t *testing.T) {

		assert.NoError(t, err)
		assert.NotEmpty(t, jwtToken)
	})

	t.Run("it should valid validate token", func(t *testing.T) {
		data, errValid := jwtInstance.Verify(jwtToken)

		assert.NoError(t, errValid)
		assert.NotNil(t, data)
	})
}
