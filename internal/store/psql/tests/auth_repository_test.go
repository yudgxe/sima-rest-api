package psql_test

import (
	"testing"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store/psql"
)

func TestAuthRepoGetUser(t *testing.T) {
	store, delete := psql.GetTestStore(t, databaseURL)
	defer delete("users")

	user := model.GetTestUser(t)
	store.User().CreateWithPermission(user, "admin")

	_, err := store.Auth().GetUser(user.Login, user.Password)
	if err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}
