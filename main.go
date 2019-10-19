package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

		e.GET("/", deadEnd)
		e.GET("/test", test)
		//e.GET("/api/travels", func(c echo.Context) error { return travels.GetTravels(c, db) })

		e.Logger.Fatal(e.Start(":7000"))

		return nil
	}

	return app
}

func deadEnd(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "It works!")
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
