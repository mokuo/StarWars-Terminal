package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
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

	// tmp ディレクトリ作成
	tmpDir, tmpDirErr := ioutil.TempDir("", "homebrew-starwars-terminal")
	if tmpDirErr != nil {
		log.Fatal(tmpDirErr)
	}
	defer os.RemoveAll(tmpDir)

	// clone
	repo, cloneErr := git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:      "https://github.com/mokuo/homebrew-starwars-terminal",
		Progress: os.Stdout,
	})
	if cloneErr != nil {
		log.Fatal(cloneErr)
	}

	// url.txt
	url := []byte(*release.TarballURL)
	urlFilePath := filepath.Join(tmpDir, "url.txt")
	urlErr := ioutil.WriteFile(urlFilePath, url, 0644)
	if urlErr != nil {
		log.Fatal(urlErr)
	}

	// sha256
	res, httpErr := http.Get(*release.TarballURL)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	defer res.Body.Close()

	hash := sha256.New()
	_, copyErr := io.Copy(hash, res.Body)
	if copyErr != nil {
		log.Fatal(copyErr)
	}

	sha256 := fmt.Sprintf("%x", hash.Sum(nil))
	sha256FilePath := filepath.Join(tmpDir, "sha256.txt")
	sha256Err := ioutil.WriteFile(sha256FilePath, []byte(sha256), 0644)
	if sha256Err != nil {
		log.Fatal(sha256Err)
	}

	// git add
	wt, wtErr := repo.Worktree()
	if wtErr != nil {
		log.Fatal(wtErr)
	}

	_, add1Err := wt.Add("url.txt")
	if add1Err != nil {
		log.Fatal(add1Err)
	}

	_, add2Err := wt.Add("sha256.txt")
	if add2Err != nil {
		log.Fatal(add2Err)
	}

	// git commit
	author := &object.Signature{
		Name:  "mokuo",
		Email: "tennis10988.yk@gmail.com",
		When:  time.Now(),
	}
	_, commitErr := wt.Commit(version, &git.CommitOptions{
		Author: author,
	})
	if commitErr != nil {
		log.Fatal(commitErr)
	}

	// git push
	pushErr := repo.Push(&git.PushOptions{
		Auth: &gitHttp.BasicAuth{
			Username: "mokuo",
			Password: os.Getenv("GITHUB_TOKEN"),
		},
	})
	if pushErr != nil {
		log.Fatal(pushErr)
	}
}
