package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shortUrl/ICAssignment/config"
)

func RegisterHandlers(router *mux.Router, config *config.Config) {

	service := NewShortnerService(config)

	urlShortRouter := router.PathPrefix("/short/v1").Subrouter()

	urlShortEndpoints(urlShortRouter, service)
}

func urlShortEndpoints(urlShortRouter *mux.Router, service *UrlService) {
	urlShortRouter.HandleFunc("/health", service.HealthCheck).Methods(http.MethodGet)
}
