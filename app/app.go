package app

import (
	"git-visualizer/app/routes"
	"log"
)

type app interface {
	startApp()
}

type App string

func (app *App) StartApp() {
	log.Println("Your Application Route Initiated")
	routes.Routes()
}
