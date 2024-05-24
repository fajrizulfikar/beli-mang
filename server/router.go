package server

import (
	"beli-mang/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (s *Server) RegisterRoute(cfg *config.Config) {
	mainRoute := s.app.Group("/v1")

	registerImageRoute(mainRoute, cfg, s.logger)
}

func registerImageRoute(e *echo.Group, cfg *config.Config, logger *zap.Logger) {
	ctr := controller.NewImageController(service.NewImageService(cfg, logger))
	auth := middleware.Authentication(cfg.JWTSecret)
	e.POST("/image", auth(ctr.PostImage))
}
