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
	exe, exeErr := os.Executable()
	if exeErr != nil {
		log.Fatal(exeErr)
	}

	exeDir := filepath.Dir(exe)
	imgDirPath := filepath.Join(exeDir, "images")

	files, err := ioutil.ReadDir(imgDirPath)
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
	exe, exeErr := os.Executable()
	if exeErr != nil {
		log.Fatal(exeErr)
	}

	exeDir := filepath.Dir(exe)
	relPath := filepath.Join(exeDir, "images", charImgFileName)
	imgFilePath, absErr := filepath.Abs(relPath)
	if absErr != nil {
		log.Fatal(absErr)
	}

	return imgFilePath
}
