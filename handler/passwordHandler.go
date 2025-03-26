package handler

import (
	"cs-p2p-sell/view/password"

	"github.com/labstack/echo/v4"
)

type PasswordHandler struct{}

func (*PasswordHandler) SnowPassword(ctx echo.Context) error {
	passwordSet := true
	if passwordSet {
		return render(ctx, password.ShowEnterPasswordPage())
	} else {
		return render(ctx, password.ShowCreatePasswordPage())
	}
}

func (*PasswordHandler) ValidatePassword(ctx echo.Context) error {
	return render(ctx, password.ShowEnterPassword())
}
