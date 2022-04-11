package test

import (
	"go_api/api/examples/jwt"
	"testing"
)

func Test(t *testing.T) {
	jwt.GetToken(t)
}

func TestRun(t *testing.T) {
	jwt.Create_JWT()
}
