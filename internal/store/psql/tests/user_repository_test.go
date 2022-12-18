package psql_test

import (
	"testing"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store/psql"
)

func TestUserRepoCreateWithPermission(t *testing.T) {
	store, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users", "privileges")

	user := model.GetTestUser(t)

	if err := store.User().CreateWithPermission(user, "admin"); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestUserRepoCreate(t *testing.T) {
	store, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	user := model.GetTestUser(t)

	if err := store.User().Create(user); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestUserRepoUpdateByLogin(t *testing.T) {
	store, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	user := model.GetTestUser(t)

	store.User().Create(user)

	user.Name = "testname"
	user.Surname = "testsurname"

	if err := store.User().UpdateByLogin(user, user.Login); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestUserRepoFindByLogin(t *testing.T) {
	store, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	user := model.GetTestUser(t)

	store.User().Create(user)

	if _, err := store.User().FindByLogin(user.Login); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}

func TestUserRepoDeleteByLogin(t *testing.T) {
	store, truncate := psql.GetTestStore(t, databaseURL)
	defer truncate("users")

	user := model.GetTestUser(t)

	store.User().CreateWithPermission(user, "admin")
	if err := store.User().DeleteByLogin(user.Login); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}
