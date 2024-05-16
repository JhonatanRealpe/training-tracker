package main

import (
	_ "github.com/JhonatanRealpe/training-tracker/cmd/api/docs"
	"github.com/JhonatanRealpe/training-tracker/infrastructure/application"
)

// @title Football Training Tracker API
// @version         1.0
// @description     This API allows users to create new training sessions for football teams, save player information, and retrieve team information. It provides endpoints for managing trainings, teams, and players.
// @host      localhost:8080
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	application.NewApp()
}
