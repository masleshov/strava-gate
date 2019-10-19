package config

import "github.com/urfave/cli"

type ConfigSet struct {
	DatabaseUrl string
	Release     bool
}

var Vars ConfigSet

func GetAppFlags(config *ConfigSet) []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "database-url",
			Value:       "come connection string",
			Usage:       "Databaes URL",
			Destination: &config.DatabaseUrl,
			EnvVar:      "DATABASE_URL",
		},
		cli.BoolFlag{
			Name:        "release",
			Usage:       "Is release mode",
			Destination: &config.Release,
			EnvVar:      "RELEASE",
		},
	}
}
