package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/mokuo/starwars-terminal/starwars/terminal"
	"github.com/urfave/cli/v2"
)

// Terminal Terminal interface.
type Terminal interface {
	Setup()
	Cmd() string
	Args() []string
}

func CharFileList() []os.FileInfo {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal(wdErr)
	}

	imgDirPath := filepath.Join(wd, "images")

	files, err := ioutil.ReadDir(imgDirPath)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// RandomCharFileName Return random character name.
func RandomCharFileName() string {
	files := CharFileList()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(files))

	return files[i].Name()
}

// ImgFilePath Return file path by character image file name.
func ImgFilePath(charImgFileName string) string {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal(wdErr)
	}

	relPath := filepath.Join(wd, "images", charImgFileName)
	imgFilePath, absErr := filepath.Abs(relPath)
	if absErr != nil {
		log.Fatal(absErr)
	}

	return imgFilePath
}

func main() {
	app := cli.NewApp()
	app.Name = "StarWars Terminal"
	app.Usage = "May the Force be with you."
	app.Action = func(c *cli.Context) error {
		var fileName string

		firstArg := c.Args().Get(0)
		if firstArg == "" {
			fileName = RandomCharFileName()
		} else {
			fileName = firstArg + ".png"
		}

		imgFilePath := ImgFilePath(fileName)

		terminal := terminal.NewIterm2()
		terminal.Setup(imgFilePath)

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
