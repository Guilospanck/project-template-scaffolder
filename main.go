package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.0-rc"
	app.Usage = "Generate scaffold project layout for Go and React web."
	app.Commands = []cli.Command{
		{
			Name:    "golang",
			Aliases: []string{"go"},
			Usage:   " Generate scaffold project layout for Go web",
			Action: func(c *cli.Context) error {
				cmd := exec.Command("make", "go-web")
				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}

				return err
			},
		},
		{
			Name:    "react",
			Aliases: []string{"react"},
			Usage:   " Generate scaffold project layout for React web",
			Action: func(c *cli.Context) error {
				cmd := exec.Command("make", "react-web-ts")
				err := cmd.Run()

				if err != nil {
					log.Fatal(err)
				}

				return err
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
