package network

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

var site = "https://api.bugfix.es"
var path = "/v1/bug"
var url = site + path

func Emit(levelno int, level, message string, args ...interface{}) {
	apiSecret := os.Getenv("BUGFIXES_SECRET")
	apiKey := os.Getenv("BUGFIXES_KEY")
	apiId := os.Getenv("BUGFIXES_ID")

	if apiSecret == "" {
		return
	}

	if apiKey == "" {
		return
	}

	if apiId == "" {
		return
	}

	dataString := fmt.Sprintf(`{"message": "%s", "loglevel": "%d"}`, message, levelno)
	jsonString := []byte(dataString)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonString))
	req.Header.Set(`Content-Type`, `application/json`)
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-ID", apiId)
	if err != nil {
		log.Fatal("Send Error", err)
	}
}
