package teststore_test

import (
	"testing"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store/teststore"
)

func TestAuthRepoGetUser(t *testing.T) {
	ts := teststore.New()
	user := model.GetTestUser(t)
	permission := "admin"

	ts.User().CreateWithPermission(user, permission)
	p, err := ts.Auth().GetUser(user.Login, user.Password)
	if err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	if p.Permission != permission {
		t.Errorf("\nIncorrect result, Expect\n%s\n, got\n%s", permission, p.Permission)
	}
}
