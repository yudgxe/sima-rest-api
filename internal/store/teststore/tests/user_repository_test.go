package teststore_test

import (
	"testing"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/store/teststore"
)

func TestUserRepoFindByLogin(t *testing.T) {
	ts := teststore.New()
	in := model.GetTestUser(t)
	ts.User().Create(in)

	out, err := ts.User().FindByLogin(in.Login)
	if err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}

	if in.Login != out.Login {
		t.Errorf("\nIncorrect result, Expect\n%v\n, got\n%v", in, out)
	}
}
