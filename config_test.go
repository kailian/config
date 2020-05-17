package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	Config, err := New(&Options{".", "test"}).ReadInConfig()
	if err != nil {
		t.Error(err)
	} else {
		VConfig := Config.GetConfig()
		assert.Equal(t, 6379, VConfig.GetInt(`redis_conn.port`))
		assert.Equal(t, "localhost", VConfig.GetString(`redis_conn.host`))
	}
}
