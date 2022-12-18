package basic_test

import (
	"testing"
	"time"

	"github.com/yudgxe/sima-rest-api/internal/model"
	"github.com/yudgxe/sima-rest-api/internal/service/basic"
	"github.com/yudgxe/sima-rest-api/internal/store/teststore"
)

func TestAuthServiceGenerateToken(t *testing.T) {
	ts := teststore.New()
	user := model.GetTestUser(t)

	ts.User().CreateWithPermission(user, "admin")

	as := basic.NewAuthService(ts, "secret", 1*time.Second)
	if _, err := as.GenerateToken(user.Login, user.Password); err != nil {
		t.Errorf("\nGot error  -> %s", err.Error())
	}
}
