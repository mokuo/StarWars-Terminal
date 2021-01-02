package util

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func CharFileList() []os.FileInfo {
	files, err := ioutil.ReadDir(imgdir())
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func RandomCharFileName() string {
	files := CharFileList()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(files))

	return files[i].Name()
}

func ImgFilePath(charImgFileName string) string {
	relPath := filepath.Join(imgdir(), charImgFileName)
	imgFilePath, absErr := filepath.Abs(relPath)
	if absErr != nil {
		log.Fatal(absErr)
	}

	return imgFilePath
}

func imgdir() string {
	exe, exeErr := os.Executable()
	if exeErr != nil {
		log.Fatal(exeErr)
	}

	sym, symErr := filepath.EvalSymlinks(exe)
	if symErr != nil {
		log.Fatal(symErr)
	}

	symDir := filepath.Dir(sym)
	return filepath.Join(symDir, "images")
}
