package server

import (
	"encoding/base64"
	"github.com/achjailani/kickoff-go-rest/config"
	"github.com/achjailani/kickoff-go-rest/pkg/security/jwt"
)

// NewJWT is constructor
func NewJWT(conf *config.Config) (*jwt.JWT, error) {
	privateKey, errPri := base64.StdEncoding.DecodeString(conf.AppPrivateKey)
	publicKey, errPub := base64.StdEncoding.DecodeString(conf.AppPublicKey)

	if errPub != nil {
		return nil, errPri
	}

	if errPub != nil {
		return nil, errPub
	}

	jwt := jwt.NewJWT(privateKey, publicKey)

	return jwt, nil
}
