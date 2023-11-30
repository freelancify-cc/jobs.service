package test

import (
	"os"
	"testing"

	"github.com/freelancify/jobs/config"
)

func TestConfig(t *testing.T) {
	t.Logf("environment: %s\n", config.GetConfig().Environment)
	t.Logf("port: %s\n", config.GetConfig().Port)
	t.Logf("host: %s\n", config.GetConfig().Host)
	t.Logf("db string: %s\n", config.GetConfig().DbString)
}
