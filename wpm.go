package main

import (
	cli "github.com/urfave/cli/v2"
	"log"
	"os"
)

/* Handle the user interaction part */
func CLIHandler() {
	app := &cli.App{
		Name: "wpm",
		Version: "v0.1",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Charles SENGES",
				Email: "charles.senges@protonmail.com",
			},
		},
		Usage: "A comprehensive tool for advanced wordpress management",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "init wpm project structure",
				Action:  WpmInit,
			},
			{
				Name:    "deploy",
				Aliases: []string{"p"},
				Usage:   "deploy changes to target environment",
				Action:  WpmDeploy,
			},
			{
				Name:    "save",
				Aliases: []string{"s"},
				Usage:   "save changes to remote vcs",
				Action:  WpmSave,
			},
			{
				Name:    "db",
				Usage:   "database related commands",
				Subcommands: []*cli.Command{
					{
						Name:  "backup",
						Usage: "create a local database backup",
						Action: WpmDbBackup,
					},
					{
						Name:  "pull",
						Usage: "pull remote database",
						Action: WpmDbPull,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/* -- Available commands -- */

/* Not Yet Implemented */
func WpmInit(c *cli.Context) error {
	return nil
}

/* Not Yet Implemented */
func WpmDeploy(c *cli.Context) error {
	return nil
}

/* Not Yet Implemented */
func WpmSave(c *cli.Context) error {
	return nil
}

/* Not Yet Implemented */
func WpmDbBackup(c *cli.Context) error {
	return nil
}

/* Not Yet Implemented */
func WpmDbPull(c *cli.Context) error {
	return nil
}