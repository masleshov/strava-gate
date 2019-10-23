package v1

import (
	"net/http"

	"bitbucket.org/virtualtrainer/strava-gate/config"
	"bitbucket.org/virtualtrainer/strava-gate/network"
	"github.com/labstack/echo/v4"
)

const verifyToken string = "ba5943e3-e0bc-49b5-9199-791bca96be05"

type callbackPostRequest struct {
	ObjectType, AspectType       string
	ObjectID, OwnerID, EventType int64
	SubscriptionID               int
}

type callbackGetRequest struct {
	Mode, Token, Challenge string
}

// CallbackPostHandler should be called by Strava when data was changed
func CallbackPostHandler(c echo.Context) error {
	req := new(callbackPostRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	return c.String(http.StatusOK, "")
}

// CallbackGetHandler should be called by Strava for checking accessibility of subscription initiator
func CallbackGetHandler(c echo.Context) error {
	req := new(callbackGetRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	if req.Mode != "subscribe" {
		echo.NewHTTPError(http.StatusForbidden, "mode is not \"subscribe\"")
	}

	if req.Token != verifyToken {
		echo.NewHTTPError(http.StatusForbidden, "invalid token. Received "+req.Token)
	}

	return c.String(http.StatusOK, "{\"hub.challenge\":"+req.Challenge+"}")
}

// Subscribe initiates a subscription to Strava on some events
func Subscribe() error {
	params := make(map[string]string)
	params["client_id"] = config.Vars.ClientID
	params["client_secret"] = config.Vars.ClientSecret
	params["callback_url"] = config.Vars.WebhookURL
	params["verify_token"] = verifyToken

	_, err := network.SendPostRequest("https://api.strava.com/api/v3/push_subscriptions", params)
	if err != nil {
		return err
	}

	return nil
}
