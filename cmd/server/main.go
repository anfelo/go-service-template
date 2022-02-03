package main

import (
	"net/http"

	"github.com/anfelo/gotodo/internal/todos"
	transportHTTP "github.com/anfelo/gotodo/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		}).Info("Setting up application")

	todosService := todos.NewService()

	handler := transportHTTP.NewHandler(todosService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	log.Info("Go Todo App")
	app := App{
		Name:    "Todo App",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting up our Web App")
		log.Fatal(err)
	}
}
