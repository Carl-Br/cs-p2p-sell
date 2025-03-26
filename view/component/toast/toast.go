package toast

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

type level string

const (
	Info    level = "info"
	Success level = "success"
	Warning level = "warning"
	Error   level = "error"
)

type toast struct {
	Level   level  `json:"level"`
	Message string `json:"message"`
}

func NewToast(level level, message string) *toast {
	return &toast{Level: level, Message: message}
}

// Send sends one or more toasts as an HX-Trigger event.
func Send(ctx echo.Context, toasts ...toast) {
	event := map[string][]toast{"makeToast": toasts}
	jsonData, _ := json.Marshal(event)
	ctx.Response().Header().Set("HX-Trigger", string(jsonData))
}
