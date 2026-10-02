package main

import (
	"os"

	"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/config"
	"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/internal/client"
	"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/internal/mcp"
	"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/internal/tools/cart"
	//"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/internal/tools/orders"
	"github.com/GangaRamPrasad2004/mcp-server-Ecommerce/internal/tools/products"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetOutput(os.Stderr) // Use stderr for logs (stdout is for JSON-RPC)
	logger.SetFormatter(&logrus.JSONFormatter{})

	c, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		logger.WithError(err).Warn("Invalid log level, using info")
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)

	logger.WithFields(logrus.Fields{
		"api_url":               c.APIURL,
		"log_level":             c.LogLevel,
		"auth_token_configured": c.AuthToken != "",
	}).Info("Starting MCP E-commerce Server")

	restClient := client.NewRestClient(c.APIURL, c.AuthToken, logger)

	toolRegistry := mcp.NewRegistry(logger)

	products.NewProductToolset(toolRegistry, restClient, logger)
	cart.NewCartToolset(toolRegistry, restClient, logger)
	//orders.NewOrderToolset(toolRegistry, restClient, logger)

	logger.WithField("tool_count", len(toolRegistry.ListTools())).Info("Registered tools")

	mcpServer := mcp.NewServer(toolRegistry, logger)

	logger.Info("Starting server on stdio transport")

	if err := mcpServer.Start(); err != nil {
		logger.WithError(err).Fatal("Server error")
	}
}
