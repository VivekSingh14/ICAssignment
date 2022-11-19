package controller

import (
	"fmt"
	"net/http"
)

func (s *UrlService) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Health is good.")
}

func (s *UrlService) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	//test case: first it will check in file if this url is already or there.
	// if no then it will insert otherwise return generic msg that it is already created or return result which is already there.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is created.")
}

func (s *UrlService) FetchShortUrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is fetched.")
}
