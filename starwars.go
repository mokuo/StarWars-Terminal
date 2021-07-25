package starwars

import (
	"fmt"
	"log"
	"os"

	"github.com/mokuo/starwars-terminal/character"
	"github.com/mokuo/starwars-terminal/terminal"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:    "StarWars Terminal",
		Usage:   "May the Force be with you.",
		Version: VERSION,
		Action: func(c *cli.Context) error {
			firstArg := c.Args().Get(0)

			terminal := terminal.NewIterm2()
			return terminal.Starwars(firstArg)
		},
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "return character list",
				Action: func(c *cli.Context) error {
					return list()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func list() error {
	charNames := character.CharNames()
	for i := 0; i < len(charNames); i++ {
		fmt.Println(charNames[i])
	}

	return nil
}
