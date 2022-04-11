package test

import (
	"go_api/api/examples/jwt"
	"testing"
)

// https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go
func Test(t *testing.T) {
	jwt.GetToken(t)
}

func TestRun(t *testing.T) {
	jwt.Create_JWT()
}
