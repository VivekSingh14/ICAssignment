package controller

import (
	"github.com/shortUrl/ICAssignment/config"
)

type UrlService struct {
	Config *config.Config
}

func NewShortnerService(config *config.Config) *UrlService {
	service := UrlService{Config: config}
	return &service
}
