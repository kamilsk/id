package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/usefathom/fathom/pkg/config"
	"github.com/usefathom/fathom/pkg/datastore"
)

type App struct {
	*cli.App
	database datastore.Datastore
	config   *config.Config
}

var app *App

func main() {
	// force all times in UTC, regardless of server timezone
	os.Setenv("TZ", "")

	// setup CLI app
	app = &App{cli.NewApp(), nil, nil}
	app.Name = "Fathom"
	app.Usage = "simple & transparent website analytics"
	app.Version = "1.0.0"
	app.HelpName = "fathom"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: ".env",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Before = before
	app.After = after
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the fathom web server",
			Action:  server,
			Flags: []cli.Flag{
				cli.StringFlag{
					EnvVar: "FATHOM_SERVER_ADDR",
					Name:   "addr,port",
					Usage:  "server address",
					Value:  ":8080",
				},
				cli.BoolFlag{
					EnvVar: "FATHOM_DEBUG",
					Name:   "debug, d",
				},
			},
		},
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "register a new admin user",
			Action:  register,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "email, e",
					Usage: "user email",
				},
				cli.StringFlag{
					Name:  "password, p",
					Usage: "user password",
				},
			},
		},
	}

	if len(os.Args) < 2 || os.Args[1] != "--version" {
		log.Infof("%s %s", app.Name, app.Version)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func before(c *cli.Context) error {
	configFile := c.String("config")
	app.config = config.Parse(configFile)
	app.database = datastore.New(app.config.Database)
	return nil
}

func after(c *cli.Context) error {
	app.database.Close()
	return nil
}
