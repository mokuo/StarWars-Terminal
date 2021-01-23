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
	cmd := exec.Command("lerna-changelog", "--from="+latestReleaseVersion(), "--next-version="+version)
	output, outputErr := cmd.Output()
	if outputErr != nil {
		log.Fatal(outputErr)
	}

	changelog := string(output)
	release := CreateRelease(version, changelog)

	fmt.Println("Created release: " + *release.Name)
	fmt.Println("Tarball URL: " + *release.TarballURL)

	UpdateFormula(*release.TarballURL)

	fmt.Println("Formula is updated.")
}
