package one_day_job

import (
	"net/http"
	"one-day-job/config"
	"one-day-job/pkg/logger"
)

type Adapter struct {
	cfg        *config.Adapter
	log        *logger.Logger
	httpClient *http.Client
}

func NewAdapter(cfg *config.Adapter, log *logger.Logger, httpClient *http.Client) *Adapter {
	return &Adapter{
		cfg:        cfg,
		log:        log,
		httpClient: httpClient,
	}
}
