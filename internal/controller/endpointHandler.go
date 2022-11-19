package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
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

	//reading the content from file.
	content, err := ioutil.ReadFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]string

	hashedValue := IfPresent(m, vars["url"])

	if hashedValue == "" {
		//creating short url with md5 crypto method and putting into a map
		data = MakeShortUrl(vars["url"])
	} else {
		fmt.Println("fetched from file: ", hashedValue)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ShortUrl is created.")

	}

	// converting it into []bytes
	convertedData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("new created: ", data[vars["url"]])
	//putting it into text file
	err = ioutil.WriteFile("data.txt", convertedData, 0)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is created.")

}

func (s *UrlService) FetchShortUrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShortUrl is fetched.")
}

func HashedCodeGenerator(str string) string {
	md5HashInBytes := md5.Sum([]byte(str))
	md5HashInString := hex.EncodeToString(md5HashInBytes[:])
	fmt.Println(md5HashInString)
	return md5HashInString
}

func MakeShortUrl(url string) map[string]string {
	short := HashedCodeGenerator(url)
	map1 := make(map[string]string)
	map1[url] = short
	return map1
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func IfPresent(m map[string]string, url string) string {

	for key, element := range m {
		if key == url {
			return element
		}
	}
	return ""
}
