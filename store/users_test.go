package store_test

import (
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/saleh-ghazimoradi/Async-api-in-Golang/fixtures"
	"github.com/saleh-ghazimoradi/Async-api-in-Golang/store"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserStore(t *testing.T) {
	env := fixtures.NewTestEnv(t)
	cleanup := env.SetupDb(t)
	t.Cleanup(func() {
		cleanup(t)
	})
	now := time.Now()

	userStore := store.NewUserStore(env.Db)
	user, err := userStore.CreateUser(context.Background(), "test@test.com", "testingpassword")
	require.NoError(t, err)
	require.Equal(t, "test@test.com", user.Email)
	require.NoError(t, user.ComparePassword("testingpassword"))
	require.Less(t, now.UnixNano(), user.CreatedAt.UnixNano())

	user2, err := userStore.GetUserById(context.Background(), user.Id)
	require.NoError(t, err)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Id, user2.Id)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.CreatedAt.UnixNano(), user2.CreatedAt.UnixNano())

	user2, err = userStore.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Id, user2.Id)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.CreatedAt.UnixNano(), user2.CreatedAt.UnixNano())
}
