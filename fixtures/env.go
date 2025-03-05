package fixtures

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/saleh-ghazimoradi/Async-api-in-Golang/config"
	"github.com/saleh-ghazimoradi/Async-api-in-Golang/store"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

type TestEnv struct {
	Config *config.Config
	Db     *sql.DB
}

func NewTestEnv(t *testing.T) *TestEnv {
	os.Setenv("ENV", string(config.EnvTest))
	conf, err := config.NewConfig()
	require.NoError(t, err)

	db, err := store.NewPostgresDB(conf)
	require.NoError(t, err)

	return &TestEnv{
		Config: conf,
		Db:     db,
	}
}

func (te *TestEnv) SetupDb(t *testing.T) func(t *testing.T) {
	m, err := migrate.New(
		fmt.Sprintf("file:///%s/migrations", te.Config.ProjectRoot),
		te.Config.DatabaseURL(),
	)
	require.NoError(t, err)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		require.NoError(t, err)
	}
	return te.TearDownDb
}

func (te *TestEnv) TearDownDb(t *testing.T) {
	_, err := te.Db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE",
		strings.Join([]string{"users", "refresh_token", "reports"}, ", ")))
	require.NoError(t, err)
}
