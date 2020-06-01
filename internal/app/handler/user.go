package handler

import(
	"encoding/json"
)

var userToken UserToken

//UserToken authorized user information
type UserToken struct {
	UserID      string `json:"user_id"`
	ClientID    string `json:"client_id"`
	Scope       string `json:"scope"`
}

func retrieveUser(user interface{}) {
	token, err := json.Marshal(user)
	json.Unmarshal(token, &userToken)
	catch(err, false)
}