package starwars

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mokuo/starwars-terminal/terminal"
	"github.com/mokuo/starwars-terminal/util"
	"github.com/urfave/cli/v2"
)

func starwars(arg string) error {
	var fileName string

	if arg == "" {
		fileName = util.RandomCharFileName()
	} else {
		fileName = arg + ".png"
	}

	imgFilePath := util.ImgFilePath(fileName)

	terminal := terminal.NewIterm2()
	terminal.Setup(imgFilePath)

	err := exec.Command(terminal.Cmd(), terminal.Args()...).Run()
	if err != nil {
		return err
	}
	return nil
}

func list() error {
	files := util.CharFileList()
	for i := 0; i < len(files); i++ {
		characterName := strings.Split(files[i].Name(), ".")[0]
		fmt.Println(characterName)
	}

	return nil
}

func Run() {
	app := &cli.App{
		Name:    "StarWars Terminal",
		Usage:   "May the Force be with you.",
		Version: "v0.0.2",
		Action: func(c *cli.Context) error {
			firstArg := c.Args().Get(0)

			return starwars(firstArg)
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
