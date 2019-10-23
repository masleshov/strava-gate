package v1

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/virtualtrainer/strava-gate/config"
	"bitbucket.org/virtualtrainer/strava-gate/db"
	"bitbucket.org/virtualtrainer/strava-gate/exception"
	"bitbucket.org/virtualtrainer/strava-gate/network"
	"github.com/labstack/echo/v4"
)

const authURL string = "https://www.strava.com/oauth"

type authResponse struct {
	TokenType    string     `json:"token_type"`
	RefreshToken string     `json:"refresh_token"`
	AccessToken  string     `json:"access_token"`
	ExpiresAt    int64      `json:"expires_at"`
	ExpiresIn    int64      `json:"expires_in"`
	User         stravaUser `json:"athlete"`
}

type stravaUser struct {
	ID float64 `json:"id"`
}

// AuthHandler POST(/v1/auth, "code" string)
func AuthHandler(c echo.Context) error {
	code, err := network.GetParamFromPost(c, "code")
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}

	params := make(map[string]string)
	params["client_id"] = config.Vars.ClientID
	params["client_secret"] = config.Vars.ClientSecret
	params["code"] = code
	params["grant_type"] = "authorization_code"
	response, err := network.SendPostRequest(authURL+"/token", params)
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}

	responseObj := new(authResponse)
	err = json.Unmarshal([]byte(response), &responseObj)
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}

	user := &db.User{
		StravaUserID: responseObj.User.ID,
		AccessToken:  responseObj.AccessToken,
		RefreshToken: responseObj.RefreshToken,
		ExpiresTo:    responseObj.ExpiresIn,
	}

	err = db.SaveUser(user)
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}
	return c.String(http.StatusOK, "Authorized successfully\n"+response)
}

// DeauthHandler POST(/v1/deauth, "token" string)
func DeauthHandler(c echo.Context) error {
	token, err := network.GetParamFromPost(c, "token")
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}

	params := make(map[string]string)
	params["access_token"] = token
	response, err := network.SendPostRequest(authURL+"/deauthorize", params)
	if err != nil {
		return exception.ThrowHTTP(http.StatusForbidden, err)
	}

	return c.String(http.StatusOK, "Deauthorized successfully\n"+response)
}
