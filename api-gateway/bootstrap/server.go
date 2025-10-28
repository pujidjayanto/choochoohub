package bootstrap

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pujidjayanto/choochoohub/api-gateway/api"
	"github.com/pujidjayanto/choochoohub/api-gateway/client"
	externalClientConfig "github.com/pujidjayanto/choochoohub/api-gateway/client/config"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/logger"
)

type ApplicationServer struct {
	Port string
	App  *fiber.App
}

func NewApplicationServer() (*ApplicationServer, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}

	// fiber has it's own server
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	log := logger.GetLogger()
	httpClient := httpclient.NewClient(log)
	externalClients := client.NewDependency(
		httpClient,
		externalClientConfig.NewExternalClientConfig(externalClientConfig.UserApi{
			Host: GetUserApiHost(),
			Port: GetUserApiPort(),
		}),
		log,
	)
	apis := api.NewDependency(externalClients)

	routes(app, apis, log)

	return &ApplicationServer{
		Port: GetServerPort(),
		App:  app,
	}, nil
}
