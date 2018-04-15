package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getHelloWorld(t *testing.T) {
	q := "Hello World!"
	p := getHelloWorld()
	assert.Equal(t, p, q)
}
