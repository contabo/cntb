package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type jmap map[string]interface{}

func HandleErrors(
	err error,
	httpResp *http.Response,
	info string,
) {
	if err != nil {
		apiError := jmap{}
		responseBody, _ := ioutil.ReadAll(httpResp.Body)
		err := json.Unmarshal(responseBody, &apiError)
		if err != nil {
			log.Error(fmt.Sprintf("Error while parsing response 500 - Internal Server Error, retry or contact support\n"))
			log.Debug(fmt.Sprintf("Error while parsing response: %v", err))
		} else {
			log.Error(fmt.Sprintf("Error %v: %v - %v\n", info, apiError["statusCode"], apiError["message"]))
			log.Debug(fmt.Sprintf("Full response: %v\n", httpResp))
		}
		log.Fatal("Aborting, due to errors")
	}
	log.Tracef("http response: %v", httpResp)
}
