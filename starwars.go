package starwars

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mokuo/starwars-terminal/terminal"
	"github.com/mokuo/starwars-terminal/util"
	"github.com/urfave/cli/v2"
)

func Run() {
	app := &cli.App{
		Name:    "StarWars Terminal",
		Usage:   "May the Force be with you.",
		Version: VERSION,
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

func starwars(arg string) error {
	var charName string

	if arg == "" {
		charName = randomCharName()
	} else {
		charName = arg
	}

	imgFilePath := util.ImgFilePath(charName)

	terminal := terminal.NewIterm2()
	terminal.Setup(imgFilePath)

	err := exec.Command(terminal.Cmd(), terminal.Args()...).Run()
	if err != nil {
		return err
	}
	return nil
}

func list() error {
	charNames := charNames()
	for i := 0; i < len(charNames); i++ {
		fmt.Println(charNames[i])
	}

	return nil
}

func charNames() []string {
	files, err := ioutil.ReadDir(util.ImgDirPath())
	if err != nil {
		log.Fatal(err)
	}

	charNames := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		fileName := files[i].Name()
		charNames[i] = strings.Split(fileName, ".")[0]
	}

	return charNames
}

func randomCharName() string {
	charNames := charNames()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(charNames))

	return charNames[i]
}
