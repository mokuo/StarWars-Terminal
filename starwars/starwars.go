package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	p, e := filepath.Abs("./images/r2d2.png")
	if e != nil {
		log.Fatal(e)
	}
	script := heredoc.Docf(`
		tell application "iTerm"
			tell current session of current window
				set background image to "%s"
				set transparency to 0.5
			end tell
		end tell
	`, p)

	app := cli.NewApp()
	app.Name = "StarWars Terminal"
	app.Usage = "May the Force be with you."
	app.Action = func(c *cli.Context) error {
		err := exec.Command("osascript", "-e", script).Run()
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
