package oauth2Client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/hprose/hprose-go"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func cacheToken(token *oauth2.Token) {
	serializedToken, err := hprose.Serialize(token, true)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not serialize token due to erros %v", err))
	}
	tokenCacheFileName := getCacheFile()
	tokenCacheFile, err := os.Create(tokenCacheFileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not open cache file %v due to errors %v", tokenCacheFile, err))
	}
	_, err = tokenCacheFile.Write(serializedToken)
	tokenCacheFile.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not write to cache file %v due to errors %v", tokenCacheFile, err))
	}
}

func RestoreTokenFromCache(oauth2User string) *oauth2.Token {
	tokenCacheFileName := getCacheFile()
	serializedToken, err := ioutil.ReadFile(tokenCacheFileName)
	if err != nil {
		return nil
	}

	var token oauth2.Token
	err = hprose.Unserialize(serializedToken, &token, true)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not deserialize token due to erros %v", err))
	}
	// check token expiry, take refresh token expiry date if present otherwise fall back to access token expiry
	accessTokenExpired := token.Expiry.Before(time.Now())
	// ... expiry date of refresh token
	jwtRefreshToken := strings.Split(token.RefreshToken, ".")[1]
	jwtRefreshTokenPayloadAsBytes, _ := base64.RawStdEncoding.DecodeString(jwtRefreshToken)
	type JwtPayload struct {
		Exp int64 `json:"exp"`
	}
	var jwtPayload JwtPayload
	json.Unmarshal(jwtRefreshTokenPayloadAsBytes, &jwtPayload)
	refreshTokenExpiryDate := time.Unix(jwtPayload.Exp, 0)
	refreshTokenExpired := refreshTokenExpiryDate.Before(time.Now())

	if accessTokenExpired && refreshTokenExpired {
		return nil
	}
	log.Debug(fmt.Sprintf("Token from cache is not outdated, reusing. Access token valid to %v. Refresh token valid to %v.", token.Expiry.String(), refreshTokenExpiryDate.String()))

	// check if user has changed
	jwtAccessTokenPayloadData := strings.Split(token.AccessToken, ".")[1]
	jwtAccessTokenPayloadAsBytes, _ := base64.RawStdEncoding.DecodeString(jwtAccessTokenPayloadData)
	type JwtAccessTokenPayload struct {
		Preferred_username string `json:"preferred_username"`
	}
	var jwtAccessTokenPayload JwtAccessTokenPayload
	json.Unmarshal(jwtAccessTokenPayloadAsBytes, &jwtAccessTokenPayload)
	if jwtAccessTokenPayload.Preferred_username != oauth2User {
		log.Debug(fmt.Sprintf("User changed (old=%v) (new=%v), forcing new token.", jwtAccessTokenPayload.Preferred_username, oauth2User))
		return nil
	}

	return &token
}

func getCacheFile() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not determine home dir: %v", err))
	}
	tokenCacheDirName := home + "/.cache/cntb"
	err = os.MkdirAll(tokenCacheDirName, os.ModePerm)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not ensure cache folder: %v", err))
	}
	tokenCacheFileName := home + "/.cache/cntb/token"
	return tokenCacheFileName
}
