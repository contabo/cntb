package util

import (
	authClient "contabo.com/cli/cntb/oauth2Client"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func GetKeycloakId(oauth2User string) string {
	// get keycloakId from jwt Token
	jwtAccessToken := authClient.RestoreTokenFromCache(oauth2User).AccessToken
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtAccessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	if err != nil {
		log.Debug(err)
	}

	if claims["sub"] == nil {
		log.Fatal("Error in getting access token.")
	}
	keycloakId := claims["sub"]
	return keycloakId.(string)
}
