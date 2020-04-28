package config

import "github.com/urfave/cli"

type ConfigSet struct {
	DatabaseURL, ClientID, ClientSecret, WebhookURL string
	Release                                         bool
}

var Vars ConfigSet

func GetAppFlags(config *ConfigSet) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "database-url",
			Value:       "SOME_URL",
			Usage:       "Databaes URL",
			Destination: &config.DatabaseURL,
			EnvVar:      "DATABASE_URL",
		},
		cli.StringFlag{
			Name:        "webhook_url",
			Value:       "SOME_URL",
			Usage:       "Webhook URL",
			Destination: &config.WebhookURL,
			EnvVar:      "WEBHOOK_URL",
		},
		cli.StringFlag{
			Name:        "client_id",
			Value:       "SOME_CLIEND_ID",
			Usage:       "Client ID",
			Destination: &config.ClientID,
			EnvVar:      "CLIENT_ID",
		},
		cli.StringFlag{
			Name:        "client_secret",
			Value:       "SOME_SECRET",
			Usage:       "Client Secret",
			Destination: &config.ClientSecret,
			EnvVar:      "CLIENT_SECRET",
		},
		cli.BoolFlag{
			Name:        "release",
			Usage:       "Is release mode",
			Destination: &config.Release,
			EnvVar:      "RELEASE",
		},
	}
}
