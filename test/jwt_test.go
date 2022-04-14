package test

import (
	jwt2 "go_api/examples/jwt"
	"testing"
)

// https://github.com/dgrijalva/jwt-go/blob/master/http_example_test.go
func Test(t *testing.T) {
	jwt2.GetToken(t)
}

func TestRun(t *testing.T) {
	jwt2.Create_JWT()
}
