package main

import (
	"os"

	"github.com/MaximeHeckel/tutum-machine/lib"
	"github.com/codegangsta/cli"
	"github.com/tutumcloud/go-tutum/tutum"
)

func main() {

	tutum.User = os.Getenv("TUTUM_USER")
	tutum.ApiKey = os.Getenv("TUTUM_APIKEY")

	app := cli.NewApp()
	app.Name = "Tutum Machine"
	app.Usage = "Provision nodes on Tutum in a single YAML file"
	app.Version = "0.0.1"

	var file string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file",
			Usage:       "select a specific tutum-machine yaml file",
			Destination: &file,
		},
	}

	app.Action = func(c *cli.Context) {
		if len(c.Args()) > 0 {
			if c.Args()[0] == "up" {
				println("Launching Tutum Machines with command:", c.Args()[0])
				if file != "" {
					println("==> Provisioned with", file)
					lib.ReadFile(file)
				} else {
					println("==> Provisioned with tutum-machine.yml")
					lib.ReadFile("tutum-machine.yml")
				}
			}
			if c.Args()[0] == "stop" {
				println("Stopping Tutum Machines with command:", c.Args()[0])
			}
			if c.Args()[0] == "rm" {
				println("Terminating Tutum Machines with command:", c.Args()[0])
			}
		} else {
			println("Check available commands with --help")
		}
	}

	app.Run(os.Args)
}
