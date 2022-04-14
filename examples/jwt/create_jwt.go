package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"go_api/internal/models"
	"time"
)

const (
	//ssh-keygen -t rsa -b 2048 -m PEM -f jwtRS256.key -q -N ""
	privKeyPath = "\\keys\\jwtRS256.key"     // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "\\keys\\jwtRS256.key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

func createToken(user string) (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"JWT",
		models.User{user, 20},
	}
	return t.SignedString(signKey)
}

func Create_JWT() string {

	signKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	verifyKey = &signKey.PublicKey
	//path, err := os.Getwd()
	//signBytes, err := ioutil.ReadFile(path + privKeyPath)
	//fatal(err)
	//
	//signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	//fatal(err)

	//verifyBytes, err := ioutil.ReadFile(path + pubKeyPath)
	//fmt.Println(path + pubKeyPath)
	//fatal(err)
	//
	//verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	//fmt.Println(verifyKey)
	//fatal(err)
	serverPort = 8080 //??
	// Make a sample token
	// In a real world situation, this token will have been acquired from
	// some other API call (see Example_getTokenViaHTTP)
	token, err := createToken("kim")
	//fmt.Printf("test token : %s, err : %s", token, err)
	fatal(err)

	return token
}
