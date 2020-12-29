package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/mokuo/starwars-terminal/starwars/terminal"
	"github.com/urfave/cli/v2"
)

// Terminal インターフェース
type Terminal interface {
	Setup()
	Cmd() string
	Args() []string
}

func main() {
	terminal := terminal.NewIterm2()
	terminal.Setup()

	app := cli.NewApp()
	app.Name = "StarWars Terminal"
	app.Usage = "May the Force be with you."
	app.Action = func(c *cli.Context) error {
		err := exec.Command(terminal.Cmd(), terminal.Args()...).Run()
		if err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
