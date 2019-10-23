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
			Value:       "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
			Usage:       "Databaes URL",
			Destination: &config.DatabaseURL,
			EnvVar:      "DATABASE_URL",
		},
		cli.StringFlag{
			Name:        "webhook_url",
			Value:       "come connection string",
			Usage:       "Webhook URL",
			Destination: &config.WebhookURL,
			EnvVar:      "WEBHOOK_URL",
		},
		cli.StringFlag{
			Name:        "client_id",
			Value:       "40068",
			Usage:       "Client ID",
			Destination: &config.ClientID,
			EnvVar:      "CLIENT_ID",
		},
		cli.StringFlag{
			Name:        "client_secret",
			Value:       "03eb67f0eac3d029174e7aa54b9d136c44974351",
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
