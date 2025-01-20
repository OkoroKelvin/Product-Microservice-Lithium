package main

import (
	"fmt"
	"subscription_service/internal/configs"
	"subscription_service/internal/handlers"
	"subscription_service/internal/repositories"
	"subscription_service/internal/services"

	"go.uber.org/dig"
)

func Bootstrap() (*ApplicationServer, error) {
	c := dig.New()

	serviceConstructors := []interface{}{
		// APPLICATION SERVER
		NewApplicationServer,

		// CONFIGS
		configs.SetUpDatabaseConnection,
		configs.NewLogger,

		// HANDLERS
		handlers.NewGRPCHandler,

		// SERVICES
		services.NewSubscriptionService,

		// REPOSITORIES
		repositories.NewSubscriptionRepository,
	}

	for _, constructor := range serviceConstructors {
		if err := c.Provide(constructor); err != nil {
			panic(fmt.Sprintf("Failed to provide service: %v", err))
		}
	}

	var app *ApplicationServer
	if err := c.Invoke(func(a *ApplicationServer) {
		app = a
	}); err != nil {
		return nil, err
	}

	return app, nil
}
