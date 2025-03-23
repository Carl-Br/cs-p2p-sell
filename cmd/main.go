package main

import (
	"cs-p2p-sell/handler"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Static("/static", "static") // for static files

	homeHandler := handler.HomeHandler{}
	passwordHandler := handler.PasswordHandler{}

	app.GET("/", homeHandler.ShowHome)
	app.GET("/password", passwordHandler.SnowPassword)
	app.POST("/password", passwordHandler.ValidatePassword)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
		slog.Info("PORT has not been set in the env, using default Port", "Port", port)
	}
	slog.Info("App running", "port", port)
	app.Logger.Fatal(app.Start(port))
}
