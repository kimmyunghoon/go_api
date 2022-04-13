package test

import (
	"go_api/api/driver"
	"testing"
)

func TestDB(t *testing.T) {
	driver.FirestoreInit(t)
}
