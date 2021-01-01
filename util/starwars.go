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
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal(wdErr)
	}

	imgDirPath := filepath.Join(wd, "images")

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
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal(wdErr)
	}

	relPath := filepath.Join(wd, "images", charImgFileName)
	imgFilePath, absErr := filepath.Abs(relPath)
	if absErr != nil {
		log.Fatal(absErr)
	}

	return imgFilePath
}
