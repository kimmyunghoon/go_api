package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go_api/api/models"
	"time"
)

func createToken(user string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"test",
		models.User{user, 20},
	}
	fmt.Println(t.Header)
	fmt.Println(t.Claims)
	return t.SignedString(signKey)
}

func Create_JWT() string {

	signKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	verifyKey = &signKey.PublicKey
	fmt.Println("verifyKey : ", verifyKey)
	serverPort = 8080 //??
	// Make a sample token
	// In a real world situation, this token will have been acquired from
	// some other API call (see Example_getTokenViaHTTP)
	token, err := createToken("kim")
	//fmt.Printf("test token : %s, err : %s", token, err)
	fatal(err)

	return token
}
