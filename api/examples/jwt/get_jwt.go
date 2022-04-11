package jwt

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func GetToken(t *testing.T) {
	serverPort = 8080 //??

	// See func authHandler for an example auth handler that produces a token
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/authenticate", serverPort), url.Values{
		"name": {"kim"},
	})
	if err != nil {
		//fatal(err)
		t.Error(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token out of the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	tokenString := strings.TrimSpace(buf.String())

	//fmt.Println(tokenString)
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify

		fmt.Println(token.Method)
		fmt.Println(token.Header)
		fmt.Println(token.Claims.(*CustomClaimsExample).User.Name)

		fmt.Println(token.Valid)
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		fmt.Println(verifyKey)
		return verifyKey, nil
	})
	fmt.Println(token, err)
	if err != nil {
		//fatal(err)
		t.Error(err)
	}

	claims := token.Claims.(*CustomClaimsExample)
	fmt.Println(claims.User.Name)

}
