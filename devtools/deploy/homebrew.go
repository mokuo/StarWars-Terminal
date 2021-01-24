package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mokuo/starwars-terminal"
)

const username string = "mokuo"

// UpdateFormula Update formula of Homebrew.
func UpdateFormula(url string) {
	workDir := tmpDir()
	defer os.RemoveAll(workDir)

	repo := gitClone(workDir)
	updateURLFile(workDir, url)
	updateSha256File(workDir, url)

	wt, wtErr := repo.Worktree()
	if wtErr != nil {
		log.Fatal(wtErr)
	}

	gitAdd(wt, "url.txt")
	gitAdd(wt, "sha256.txt")
	gitCommit(wt)
	gitPush(repo)
}

func tmpDir() string {
	tmpDir, tmpDirErr := ioutil.TempDir("", "homebrew-starwars-terminal")
	if tmpDirErr != nil {
		log.Fatal(tmpDirErr)
	}

	return tmpDir
}

func gitClone(workDir string) *git.Repository {
	repo, cloneErr := git.PlainClone(workDir, false, &git.CloneOptions{
		URL:      "https://github.com/mokuo/homebrew-starwars-terminal",
		Progress: os.Stdout,
	})
	if cloneErr != nil {
		log.Fatal(cloneErr)
	}

	return repo
}

func updateURLFile(workDir, url string) {
	urlFilePath := filepath.Join(workDir, "url.txt")
	urlErr := ioutil.WriteFile(urlFilePath, []byte(url), 0644)
	if urlErr != nil {
		log.Fatal(urlErr)
	}
}

func updateSha256File(workDir, url string) {
	res, httpErr := http.Get(url)
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
	sha256FilePath := filepath.Join(workDir, "sha256.txt")
	sha256Err := ioutil.WriteFile(sha256FilePath, []byte(sha256), 0644)
	if sha256Err != nil {
		log.Fatal(sha256Err)
	}
}

func gitAdd(wt *git.Worktree, file string) {
	_, addErr := wt.Add(file)
	if addErr != nil {
		log.Fatal(addErr)
	}
}

func gitCommit(wt *git.Worktree) {
	author := &object.Signature{
		Name:  username,
		Email: "tennis10988.yk@gmail.com",
		When:  time.Now(),
	}
	_, commitErr := wt.Commit(starwars.Version(), &git.CommitOptions{
		Author: author,
	})
	if commitErr != nil {
		log.Fatal(commitErr)
	}
}

func gitPush(repo *git.Repository) {
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
