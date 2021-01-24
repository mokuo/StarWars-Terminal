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
	currentVersion := latestReleaseVersion()
	nextVersion := starwars.Version()
	if currentVersion == nextVersion {
		log.Fatal("Update the version of version.go file.")
	}

	cmd := exec.Command("lerna-changelog", "--from="+currentVersion, "--next-version="+nextVersion)
	output, outputErr := cmd.Output()
	if outputErr != nil {
		fmt.Println(string(output))
		log.Fatal(outputErr)
	}

	changelog := string(output)
	release := CreateRelease(nextVersion, changelog)

	fmt.Println("Created release: " + *release.Name)

	UpdateFormula(*release.TarballURL)

	fmt.Println("Formula is updated.")
}
