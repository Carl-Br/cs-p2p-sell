package handler

import (
	"cs-p2p-sell/view/home"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func (*HomeHandler) ShowHome(ctx echo.Context) error {
	return render(ctx, home.ShowHome())
}
