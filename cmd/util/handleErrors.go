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
		json.Unmarshal(responseBody, &apiError)
		log.Error(fmt.Sprintf("Error %v: %v - %v\n", info, apiError["statusCode"], apiError["message"]))
		log.Debug(fmt.Sprintf("Full response: %v\n", httpResp))
		log.Fatal("Aborting, due to errors")
	}
	log.Tracef("http response: %v", httpResp)
}
