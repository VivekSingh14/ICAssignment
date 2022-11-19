package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *UrlService) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Health is good.")
}

func (s *UrlService) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	//test case: first it will check in file if this url is already or there.
	// if no then it will insert otherwise return generic msg that it is already created or return result which is already there.
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	//result := InsertURL(collection, vars["url"])
	//fmt.Fprintf(w, "Hash: %v\n", result)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is created.")

	//val := "old falcon\n"
	data := []byte(vars["url"])

	err := ioutil.WriteFile("data.txt", data, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")

}

func (s *UrlService) FetchShortUrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is fetched.")
}
