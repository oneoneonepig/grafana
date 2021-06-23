package notifications

import (
	"testing"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/stretchr/testify/require"
)

func TestEmailCodes(t *testing.T) {
	t.Run("When generating code", func(t *testing.T) {
		setting.EmailCodeValidMinutes = 120

		user := &models.User{Id: 10, Email: "t@a.com", Login: "asd", Password: "1", Rands: "2"}
		code, err := createUserEmailCode(user, nil)
		require.NoError(t, err)

		t.Run("getLoginForCode should return login", func(t *testing.T) {
			login := getLoginForEmailCode(code)
			require.Equal(t, "asd", login)
		})

		t.Run("Can verify valid code", func(t *testing.T) {
			isValid, err := validateUserEmailCode(user, code)
			require.NoError(t, err)
			require.True(t, isValid)
		})

		t.Run("Cannot verify in-valid code", func(t *testing.T) {
			code = "ASD"
			isValid, err := validateUserEmailCode(user, code)
			require.NoError(t, err)
			So(isValid, ShouldBeFalse)
		})
	})
}
