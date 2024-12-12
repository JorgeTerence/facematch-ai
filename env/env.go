package env

import (
	"strings"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
	Testing     Environment = "tests"
)

func FromString(str string) Environment {
	switch strings.ToLower(str) {
	case "development", "dev":
		return Development
	case "production", "prod":
		return Production
	case "testing", "test":
		return Testing
	default:
		return Development
	}
}
