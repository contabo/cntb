package oauth2Client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var bearerHttpClient *http.Client = nil

func BearerHttpClient(tokenUrl string, clientId string, clientSecret string, user string, pass string) *http.Client {
	if bearerHttpClient != nil {
		return bearerHttpClient
	}
	// sanity checks
	parsedTokenUrl, err := url.ParseRequestURI(tokenUrl)
	if err != nil {
		log.Fatal("tokenURL is not a well formed URL")
	}
	if clientId == "" {
		log.Fatal("clientId is empty")
	}
	if clientSecret == "" {
		log.Fatal("clientSecret is empty")
	}
	if user == "" {
		log.Fatal("user is empty")
	}
	if pass == "" {
		log.Fatal("pass is empty")
	}

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: parsedTokenUrl.String(),
		},
	}

	// check if token has been cached
	token := RestoreTokenFromCache()

	if token == nil {
		token, err = conf.PasswordCredentialsToken(ctx, user, pass)
		log.Debug("Getting new token, as cache empty or outdated...")
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not get access token due to an error: %v", err))
		}
	}
	log.Debug(fmt.Sprintf("Got access token: %v\n", token.AccessToken))
	cacheToken(token)

	// extract tenantId and customerId
	jwt := strings.Split(token.AccessToken, ".")[1]
	jwtPayloadAsBytes, _ := base64.RawStdEncoding.DecodeString(jwt)
	type JwtPayload struct {
		TenantId   string `json:"tenantId"`
		CustomerId string `json:"customerId"`
	}
	var jwtPayload JwtPayload
	json.Unmarshal(jwtPayloadAsBytes, &jwtPayload)
	log.Debug(fmt.Sprintf("tenantId/customerId: %v/%v", jwtPayload.TenantId, jwtPayload.CustomerId))

	client := conf.Client(ctx, token)
	bearerHttpClient = client

	return bearerHttpClient
}
