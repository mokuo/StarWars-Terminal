package terminal

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
)

// Iterm2 iTerm2
type Iterm2 struct{}

// NewIterm2 コンストラクタ
func NewIterm2() Iterm2 {
	return Iterm2{}
}

// ScriptPath スクリプトのパスを返す
func ScriptPath() string {
	homeDir, err := os.UserHomeDir()
	if err != err {
		log.Fatal(err)
	}

	return homeDir + "/Library/ApplicationSupport/iTerm2/Scripts/starwars.py"
}

// Setup セットアップ
func (t Iterm2) Setup(imgPath string) {
	script := heredoc.Docf(`
		#!/usr/bin/env python3.7
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
	`, imgPath)

	err := ioutil.WriteFile(ScriptPath(), []byte(script), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Cmd 実行コマンド
func (t Iterm2) Cmd() string {
	homeDir, err := os.UserHomeDir()
	if err != err {
		log.Fatal(err)
	}

	return homeDir + "/Library/ApplicationSupport/iTerm2/iterm2env/versions/3.7.9/bin/python3"
}

// Args コマンドの引数
func (t Iterm2) Args() []string {
	return []string{ScriptPath()}
}
