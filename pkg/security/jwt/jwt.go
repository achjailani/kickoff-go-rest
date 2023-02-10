package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// ErrInvalidToken is an error indicates token invalid
	ErrInvalidToken = errors.New("error invalid token")
	// ErrExpiredToken is an error indicates token expired
	ErrExpiredToken = errors.New("error expired token")
)

// JWT is struct
type JWT struct {
	privateKey []byte
	publicKey  []byte
}

// JWTPayload is
type JWTPayload struct {
	Identifier string `json:"id"`
	Name       string `json:"name"`
}

// JWTClaim is struct
type JWTClaim struct {
	Dat       *JWTPayload `json:"dat"`
	Audience  string      `json:"aud,omitempty"`
	ExpiresAt int64       `json:"exp,omitempty"`
	IssuedAt  int64       `json:"iat,omitempty"`
	Issuer    string      `json:"iss,omitempty"`
	NotBefore int64       `json:"nbf,omitempty"`
}

// Valid check validity payload
func (pl *JWTClaim) Valid() error {
	now := time.Now().Unix()
	if now > pl.ExpiresAt {
		return ErrExpiredToken
	}

	return nil
}

// JWTOption is a type for jwt option
type JWTOption func(jwt *JWT)

// NewJWT is constructor
func NewJWT(privateKey []byte, publicKey []byte, opts ...JWTOption) *JWT {
	j := &JWT{privateKey: privateKey, publicKey: publicKey}

	for _, opt := range opts {
		opt(j)
	}

	return j
}

// Create is a function to create jwt token
func (j *JWT) Create(ttl time.Duration, payload *JWTPayload) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now()
	claims := &JWTClaim{
		Dat:       payload,
		ExpiresAt: now.Add(ttl).Unix(),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		Issuer:    "Hell Yeah Provider",
		Audience:  "*",
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

// Verify is function to validate token
func (j *JWT) Verify(token string) (*JWTPayload, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, ErrInvalidToken
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTClaim{}, keyFunc)
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		switch {
		case ok && errors.Is(vErr.Inner, ErrExpiredToken):
			return nil, ErrExpiredToken
		default:
			return nil, ErrInvalidToken
		}
	}

	claims, ok := jwtToken.Claims.(*JWTClaim)
	if ok && jwtToken.Valid {
		return claims.Dat, nil
	}

	return nil, ErrInvalidToken
}
