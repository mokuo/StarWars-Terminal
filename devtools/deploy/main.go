package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mokuo/starwars-terminal"
)

func latestReleaseVersion() string {
	release := GetLatestRelease()
	return *release.TagName
}

func main() {
	version := starwars.Version()
	cmd := exec.Command("lerna-changelog", "--from="+latestReleaseVersion())
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	changelog := string(output)

	release := CreateRelease(version, changelog)
	fmt.Println("Created release: " + *release.Name)
}
