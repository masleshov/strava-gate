package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	v1 "bitbucket.org/virtualtrainer/strava-gate/api/v1"
	"bitbucket.org/virtualtrainer/strava-gate/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	setSignalListener()

	err := createApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createApp() *cli.App {
	config.Vars = config.ConfigSet{}

	app := cli.NewApp()
	app.Name = "Strava Gate"
	app.Flags = config.GetAppFlags(&config.Vars)

	app.Action = func(c *cli.Context) error {
		bytes, _ := json.Marshal(config.Vars)
		log.Printf("Using config: %v", string(bytes))

		e := echo.New()
		e.Debug = !config.Vars.Release

		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		e.POST("/v1/auth", v1.AuthHandler)
		e.POST("/v1/deauth", v1.DeauthHandler)
		e.POST("/v1/subscribe", v1.SubscribeHandler)
		e.POST("/v1/webhook", v1.CallbackPostHandler)
		e.GET("/v1/webhook", v1.CallbackGetHandler)

		e.Logger.Fatal(e.Start(":7000"))

		return nil
	}

	return app
}

func setSignalListener() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Infof("Got signal %v. Exiting...", sig)
		os.Exit(0)
	}()
}
