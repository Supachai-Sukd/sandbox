//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue string
		osValue      string
		want         string
	}{
		{"get string from environment", "test-key", "default-value", "test-value", "test-value"},
		{"get defualt string if not config", "", "default-value", "", "default-value"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := func(key string) string {
				return map[string]string{
					tc.key: tc.osValue,
				}[key]
			}
			cfg.SetEnvGetter(mCfg)

			got := cfg.envString(tc.key, tc.defaultValue)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIntEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue int
		osValue      string
		want         int
	}{
		{"get int from environment", "test-key", 10, "1", 1},
		{"get defualt int if not config", "", 10, "", 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := func(key string) string {
				return map[string]string{
					tc.key: tc.osValue,
				}[key]
			}
			cfg.SetEnvGetter(mCfg)

			got := cfg.envInt(tc.key, tc.defaultValue)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBoolEnvironment(t *testing.T) {
	cfg := New()
	tests := []struct {
		name         string
		key          string
		defaultValue bool
		osValue      string
		want         bool
	}{
		{"get bool from environment", "test-key", true, "FALSE", false},
		{"get defualt bool if not config", "", true, "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := func(key string) string {
				return map[string]string{
					tc.key: tc.osValue,
				}[key]
			}
			cfg.SetEnvGetter(mCfg)

			got := cfg.envBool(tc.key, tc.defaultValue)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGetAllConfig(t *testing.T) {
	cfg := New()
	tests := []struct {
		name string
		mock map[string]string
		want Config
	}{
		{"config hostname env should return as changed",
			map[string]string{cHostname: "test-hostname"},
			Config{
				Server:       Server{Hostname: "test-hostname", Port: 2566},
				DBConnection: dDBConnection}},
		{"config port env should return as changed",
			map[string]string{cPort: "4444"},
			Config{
				Server:       Server{Port: 4444},
				DBConnection: dDBConnection}},

		{"config DB_CONNECTION env should return as changed",
			map[string]string{cDBConnection: "test-db-connection"},
			Config{
				Server:       Server{Port: 2566},
				DBConnection: "test-db-connection"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mCfg := func(key string) string {
				return tc.mock[key]
			}
			cfg.SetEnvGetter(mCfg)

			got := cfg.All()

			assert.Equal(t, tc.want, got)
		})
	}
}
