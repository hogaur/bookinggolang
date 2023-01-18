package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	emails := []string{"hariomgmail.com", "hariom@gmail.com"}

	assert.Equal(t, false, IsValidEmail(emails[0]))
	assert.Equal(t, true, IsValidEmail(emails[1]))
}
