package util

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleErrors(
	err error,
	httpResp *http.Response,
	info string,
) {
	if err != nil {
		log.Error(fmt.Sprintf("Error %v: %v; %v\n", info, err, httpResp))
		log.Debug(fmt.Sprintf("Full response: %v\n", httpResp))
		log.Fatal("Aborting, due to errors")
	}
	log.Tracef("http response: %v", httpResp)
}
