package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDbForTest(t *testing.T) {
	initDB("gochat")
	db := GetDb("gochat")
	assert.NotNil(t, db, "Database instance should not be nil")
}
