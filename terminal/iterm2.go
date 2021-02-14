package terminal

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/mokuo/starwars-terminal/util"
)

func NewIterm2() Iterm2 {
	return Iterm2{}
}

type Iterm2 struct{}

func (t Iterm2) Starwars(arg string) error {
	var charName string

	if arg == "" {
		charName = util.RandomCharName()
	} else {
		charName = arg
	}

	imgFilePath := util.ImgFilePath(charName)
	setup(imgFilePath)

	err := exec.Command(cmd(), args()...).Run()
	if err != nil {
		return err
	}
	return nil
}

func setup(imgFilePath string) {
	script := heredoc.Docf(`
		#!/usr/bin/env python3
		import iterm2
		
		async def main(connection):
		    app = await iterm2.async_get_app(connection)
		    # ref: https://iterm2.com/python-api/examples/setprofile.html
		    current_profile = await app.current_terminal_window.current_tab.current_session.async_get_profile()
		
		    # ref: https://iterm2.com/python-api/profile.html?highlight=backgroundimagemode#iterm2.Profile.async_set_background_image_location
		    await current_profile.async_set_background_image_location("%s")
		    # ref: https://iterm2.com/python-api/profile.html?highlight=backgroundimagemode#iterm2.BackgroundImageMode
		    await current_profile.async_set_background_image_mode(1)
		
		iterm2.run_until_complete(main)
	`, imgFilePath)

	err := ioutil.WriteFile(scriptpath(), []byte(script), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Cmd Return latest python3 path.
func cmd() string {
	pythonDirPath := filepath.Join(iterm2home(), "iterm2env", "versions")

	files, err := ioutil.ReadDir(pythonDirPath)
	if err != nil {
		log.Fatal(err)
	}

	// ~/Library/ApplicationSupport/iTerm2/iterm2env/versions/*/bin/python3
	// ref: https://iterm2.com/python-api/tutorial/running.html#command-line
	return filepath.Join(pythonDirPath, latestVersion(files), "bin", "python3")
}

// Args Return command arguments.
func args() []string {
	return []string{scriptpath()}
}

func iterm2home() string {
	homeDir, err := os.UserHomeDir()
	if err != err {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, "Library", "ApplicationSupport", "iTerm2")
}

func scriptpath() string {
	return filepath.Join(iterm2home(), "Scripts", "starwars.py")
}

func latestVersion(files []os.FileInfo) string {
	sort.Slice(files, func(i, j int) bool {
		iMajor, iMinor, iPatch := util.GetVersions(files[i].Name())
		jMajor, jMinor, jPatch := util.GetVersions(files[i].Name())

		if iMajor != jMajor {
			return iMajor > jMajor
		}
		if iMinor != jMinor {
			return iMinor > jMinor
		}
		return iPatch > jPatch
	})

	return files[0].Name()
}
