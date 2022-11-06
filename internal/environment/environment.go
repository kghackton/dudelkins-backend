package environment

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
)

type Environment struct {
	Postgres

	InsEndpoint string
}

type Postgres struct {
	Username, Password, Host, Port, Database string

	DebugQuery bool
}

func (e *Postgres) fillFromEnv() {
	e.Username = os.Getenv("POSTGRES_USER")
	e.Password = os.Getenv("POSTGRES_PASSWORD")
	e.Host = os.Getenv("POSTGRES_HOST")
	e.Port = os.Getenv("POSTGRES_PORT")
	e.Database = os.Getenv("POSTGRES_DB")

	if debugEnv, exists := os.LookupEnv("POSTGRES_DEBUG_QUERY"); exists {
		e.DebugQuery, _ = strconv.ParseBool(debugEnv)
	}
}

func (e *Postgres) FormConnStringPgWithPool() string {
	poolSize := runtime.NumCPU() * 4
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%d", e.Username, e.Password, e.Host, e.Port, e.Database, poolSize)
}

func (e *Postgres) FormConnStringPg() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", e.Username, e.Password, e.Host, e.Port, e.Database)
}

func NewEnvironment() (e Environment) {
	e.Postgres.fillFromEnv()
	e.InsEndpoint = os.Getenv("INS_ENDPOINT")

	return
}
