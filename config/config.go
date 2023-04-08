package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type env func(key string) string

type cfg struct {
	getEnv env
}

func New() *cfg {
	return &cfg{getEnv: os.Getenv}
}

type Config struct {
	Server       Server
	DBConnection string
}

type Server struct {
	Hostname string
	Port     int
}

const (
	cHostname = "HOSTNAME"
	cPort     = "PORT"

	cDBConnection = "DB_CONNECTION"
)

const (
	dPort         = 2566
	dDBConnection = "postgresql://postgres:password@localhost:5432/demo_sandbox?sslmode=disable"
)

func InitDB(db *sql.DB) error {
	// Read the SQL file
	content, err := ioutil.ReadFile("db/01-init.sql")
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	// Execute the SQL statements
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to execute SQL statements: %w", err)
	}

	return nil
}

func (c *cfg) All() Config {
	return Config{
		Server: Server{
			Hostname: c.envString(cHostname, ""),
			Port:     c.envInt(cPort, dPort),
		},
		DBConnection: c.envString(cDBConnection, dDBConnection),
	}
}

func (c *cfg) SetEnvGetter(overrideEnvGetter env) {
	c.getEnv = overrideEnvGetter
}

func (c *cfg) envString(key, defaultValue string) string {
	val := c.getEnv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func (c *cfg) envInt(key string, defaultValue int) int {
	v := c.getEnv(key)

	val, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return val
}

func (c *cfg) envBool(key string, defaultValue bool) bool {
	v := c.getEnv(key)

	val, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}

	return val
}
