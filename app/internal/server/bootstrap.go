package server

import (
	"net/http"

	"github.com/alfattd/crud/internal/platform/config"
	"github.com/alfattd/crud/internal/platform/logger"
	"github.com/alfattd/crud/internal/platform/monitor"
)

func Run() (*config.Config, *http.Server) {
	cfg := config.Load()

	logger.New()

	monitor.Init()

	srv := New(cfg)

	return cfg, srv
}
