package mocks

import (
	"context"
	"net/http"
	"strings"

	"github.com/gmbanaag/wallet2020/internal/app/config"
)

//UserToken tokeninfo struct
type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	UserID      string `json:"user_id"`
	ClientID    string `json:"client_id"`
	Scope       string `json:"scope"`
}

//MockAuth Auth middleware mock
type MockAuth struct {
	Config *config.Config
}

//Authenticate mock
func (m MockAuth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		accessToken := strings.Trim(authorization[len("Bearer"):], " ")

		userToken := UserToken{}
		ctx := context.WithValue(r.Context(), m.Config.UserCtxKey, userToken)

		if accessToken == "xxx" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if accessToken == "validfortestingpurposes" {
			userToken := UserToken{}
			userToken.AccessToken = accessToken
			userToken.UserID = r.Header.Get("X-User-ID")
			userToken.ExpiresIn = 3600
			userToken.ClientID = "cec0482b1b77d46ab7f13b114e79ae3b3c01286d"
			userToken.Scope = "default"

			ctx = context.WithValue(r.Context(), m.Config.UserCtxKey, userToken)
		} else if accessToken == "validfortestingpurposesadmin" {
			userToken := UserToken{}
			userToken.AccessToken = accessToken
			userToken.UserID = r.Header.Get("X-User-ID")
			userToken.ExpiresIn = 3600
			userToken.ClientID = "cec0482b1b77d46ab7f13b114e79ae3b3c01286d"
			userToken.Scope = "admin"
			ctx = context.WithValue(r.Context(), m.Config.UserCtxKey, userToken)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
