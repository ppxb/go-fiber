package env

import (
	"flag"
	"fmt"
	"strings"
)

var (
	active Environment
	dev    Environment = &env{value: "dev"}
	prod   Environment = &env{value: "prod"}
)

type Environment interface {
	Value() string
	IsDev() bool
	IsProd() bool
}

type env struct {
	value string
}

func (e *env) Value() string {
	return e.value
}

func (e *env) IsDev() bool {
	return e.value == "dev"
}

func (e *env) IsProd() bool {
	return e.value == "prod"
}

func init() {
	env := flag.String(
		"env",
		"",
		"chose your active environment with server start:\n dev: development environment\n prod: production environment\n",
	)
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case "dev":
		active = dev
	case "prod":
		active = prod
	default:
		active = dev
		fmt.Println("Warning: '-env' cannot be found, or it is illegal. The default environment 'dev' will be used.")
	}
}

func Active() Environment {
	return active
}
