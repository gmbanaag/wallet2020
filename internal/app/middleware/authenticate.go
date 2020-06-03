package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"net/http"
	"net/url"

	"github.com/gmbanaag/wallet2020/internal/app/config"
)

//Auth object
type Auth struct {
	Config *config.Config
}

//UserToken object
type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	UserID      string `json:"user_id"`
	ClientID    string `json:"client_id"`
	Scope       string `json:"scope"`
}

//Authenticate checks access_token validity
func (a Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userAgent := r.Header.Get("User-Agent")
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		accessToken := strings.Trim(authorization[len("Bearer"):], " ")

		userToken := UserToken{}
		ctx := context.WithValue(r.Context(), a.Config.UserCtxKey, userToken)

		if accessToken == "cec0482b1b77d46ab7f13b114e79ae3b3c01286d" {
			userToken := UserToken{}
			userToken.AccessToken = accessToken
			userToken.UserID = "ff7cc44a-b949-413c-9c75-6f34a5699915"
			userToken.ExpiresIn = 3600
			userToken.ClientID = "cec0482b1b77d46ab7f13b114e79ae3b3c01286d"
			userToken.Scope = "admin"
			ctx = context.WithValue(r.Context(), a.Config.UserCtxKey, userToken)
		} else if accessToken == "ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b" {
			userToken := UserToken{}
			userToken.AccessToken = accessToken
			userToken.UserID = "ff7cc44a-b949-413c-9c75-6f34a5699915"
			userToken.ExpiresIn = 3600
			userToken.ClientID = "cec0482b1b77d46ab7f13b114e79ae3b3c01286d"
			userToken.Scope = "default"
			ctx = context.WithValue(r.Context(), a.Config.UserCtxKey, userToken)
		} else if accessToken != "" {
			server := Server{Host: a.Config.OAuthEndpoint, TokenInfoPath: a.Config.OAuthTokeninfo, UserAgent: userAgent}
			user, valid := server.GetTokenInfo(accessToken)

			if valid == false && user.UserID == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userToken.AccessToken = user.AccessToken
			userToken.UserID = user.UserID
			userToken.ExpiresIn = user.ExpiresIn
			userToken.ClientID = user.ClientID
			ctx = context.WithValue(r.Context(), a.Config.UserCtxKey, userToken)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//TokenInfoResponse from identity service
type TokenInfoResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	UserID      string `json:"user_id"`
	ClientID    string `json:"client_id"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
}

//Server object to connect to identity service
type Server struct {
	Host          string
	TokenInfoPath string
	UserAgent     string
}

//GetTokenInfo calls /tokeninfo api for token validation
func (s Server) GetTokenInfo(accessToken string) (TokenInfoResponse, bool) {
	apiURL := s.Host

	tokenInfoResponse := TokenInfoResponse{}

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = s.TokenInfoPath
	u.RawQuery = fmt.Sprintf("access_token=%s", accessToken)
	urlStr := fmt.Sprintf("%v", u)

	resp, err := http.Get(urlStr)
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return tokenInfoResponse, false
	}
	defer resp.Body.Close()
	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tokenInfoResponse, false
	}

	err = json.NewDecoder(bytes.NewReader(bits)).Decode(&tokenInfoResponse)
	if err != nil {
		log.Println(err.Error())
		return tokenInfoResponse, false
	}

	if tokenInfoResponse.Error != "" {
		return tokenInfoResponse, false
	}

	return tokenInfoResponse, true
}
