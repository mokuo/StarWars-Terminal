package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/MakeNowJust/heredoc"
	"github.com/urfave/cli"
)

func main() {
	script := heredoc.Doc(`
		tell application "iTerm"
			activate
			tell current session of current window
				set background image to "~/Downloads/starwars-avatars/pngs/starwars-r2-d2.png"
			end tell
		end tell
	`)

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
