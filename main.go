package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func execCommands(path string) error {
	cmd := exec.Command("git", "clone", fmt.Sprintf("https://github.com/Guilospanck/%s", path))
	value, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", value)
		log.Fatal(err)
		return err
	}

	cmd = exec.Command("rm", "-rf", fmt.Sprintf("%s/.git", path))
	value, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", value)
		log.Fatal(err)
		return err
	}

	return nil
}

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
				return execCommands("go-web-template")
			},
		},
		{
			Name:    "react",
			Aliases: []string{"react"},
			Usage:   " Generate scaffold project layout for React web",
			Action: func(c *cli.Context) error {
				return execCommands("react-18-typescript")
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
