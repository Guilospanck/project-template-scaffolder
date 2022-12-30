package main

import (
	"fmt"
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
				value, err := cmd.CombinedOutput()

				if err != nil {
					fmt.Printf("%s\n", value)
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
				value, err := cmd.CombinedOutput()

				if err != nil {
					fmt.Printf("%s\n", value)
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
