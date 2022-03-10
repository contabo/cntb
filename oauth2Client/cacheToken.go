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

func RestoreTokenFromCache() *oauth2.Token {
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
	jwt := strings.Split(token.RefreshToken, ".")[1]
	jwtPayloadAsBytes, _ := base64.RawStdEncoding.DecodeString(jwt)
	type JwtPayload struct {
		Exp int64 `json:"exp"`
	}
	var jwtPayload JwtPayload
	json.Unmarshal(jwtPayloadAsBytes, &jwtPayload)
	refreshTokenExpiryDate := time.Unix(jwtPayload.Exp, 0)
	refreshTokenExpired := refreshTokenExpiryDate.Before(time.Now())

	if accessTokenExpired && refreshTokenExpired {
		return nil
	}
	log.Debug(fmt.Sprintf("Token from cache is not outdated, resusing. Access token valid to %v. Refresh token valid to %v.", token.Expiry.String(), refreshTokenExpiryDate.String()))
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
