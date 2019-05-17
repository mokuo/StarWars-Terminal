package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/MakeNowJust/heredoc"
	"github.com/urfave/cli"
)

func main() {
	gopath := os.Getenv("GOPATH")
	script := heredoc.Docf(`
		tell application "iTerm"
			tell current session of current window
				set background image to "%s/src/github.com/mokuo/StarWars-Terminal/images/r2-d2.png"
				set transparency to 0.5
			end tell
		end tell
	`, gopath)

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
