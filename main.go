package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Tutum Machine"
	app.Usage = "Provision nodes on Tutum in a single YAML file"
	app.Version = "0.0.1"

	var file string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file",
			Value:       "tutum-machine.yml",
			Usage:       "Select a specific tutum-machine yaml file",
			Destination: &file,
		},
	}

	app.Action = func(c *cli.Context) {
		if len(c.Args()) > 0 {
			if c.Args()[0] == "up" {
				println("Launching Tutum Machines with command:", c.Args()[0])
			}
			if c.Args()[0] == "stop" {
				println("Stopping Tutum Machines with command:", c.Args()[0])
			}
			if c.Args()[0] == "rm" {
				println("Terminating Tutum Machines with command:", c.Args()[0])
			}

		}

		if file != "" {
			println("Provisioned with", file)
		}
	}

	app.Run(os.Args)
}
