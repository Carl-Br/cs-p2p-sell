package handler

import (
	"cs-p2p-sell/view/component/toast"
	"cs-p2p-sell/view/pin"

	"github.com/labstack/echo/v4"
)

type PasswordHandler struct{}

func (*PasswordHandler) SnowPassword(ctx echo.Context) error {
	passwordSet := true
	if passwordSet {
		return render(ctx, pin.ShowEnterPassword())
	} else {
		return render(ctx, pin.ShowCreatePassword())
	}
}

func (*PasswordHandler) ValidatePassword(ctx echo.Context) error {
	toast.Success(ctx, "password is correct")
	return render(ctx, pin.ShowEnterPassword())
}
