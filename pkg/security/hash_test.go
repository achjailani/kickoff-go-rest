package security_test

import (
	"github.com/achjailani/kickoff-go-rest/pkg/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMake(t *testing.T) {
	password := "Hell@Yeah"

	hash := security.HashMake(password)

	t.Run("it should not be empty", func(t *testing.T) {
		assert.NotEmpty(t, hash)
	})

	t.Run("it should not be equal", func(t *testing.T) {
		assert.NotEqual(t, password, hash)
	})

	verified := security.HashVerify(password, hash)

	t.Run("it should be valid verify", func(t *testing.T) {
		assert.True(t, verified)
	})

	password2 := "Hell@Yooo"

	verified2 := security.HashVerify(password2, hash)
	t.Run("it should invalid", func(t *testing.T) {
		assert.False(t, verified2)
	})
}
