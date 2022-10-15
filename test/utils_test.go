package test

import (
	"fmt"
	"testing"
)

type H map[string]interface{}

func TestMap(t *testing.T) {
	fmt.Printf("%+v", H{
		"ping": "pong",
	})
}
