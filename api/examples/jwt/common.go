package jwt

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"go_api/api/models"
	"log"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	serverPort int
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	models.User
}
