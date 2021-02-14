package util

import (
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

// RandomCharName Return random character name.
func RandomCharName() string {
	charNames := CharNames()

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(charNames))

	return charNames[i]
}

// CharNames Return character names.
func CharNames() []string {
	files, err := ioutil.ReadDir(ImgDirPath())
	if err != nil {
		log.Fatal(err)
	}

	charNames := make([]string, len(files))
	for i := 0; i < len(files); i++ {
		fileName := files[i].Name()
		charNames[i] = strings.Split(fileName, ".")[0]
	}

	return charNames
}
