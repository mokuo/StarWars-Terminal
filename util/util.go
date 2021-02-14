package util

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ImgFileExt extention of character image file
const ImgFileExt string = "png"

// ImgFilePath Return image file path of character.
func ImgFilePath(charName string) string {
	relPath := filepath.Join(ImgDirPath(), charName+"."+ImgFileExt)
	imgFilePath, absErr := filepath.Abs(relPath)
	if absErr != nil {
		log.Fatal(absErr)
	}

	return imgFilePath
}

// ImgDirPath Return image directory path of character.
func ImgDirPath() string {
	if isTest() {
		return testImgDir()
	}

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

// GetVersions ex) "1.2.3" => 1, 2, 3
func GetVersions(s string) (int, int, int) {
	versions := strings.Split(s, ".")
	v := strs2ints(versions)
	major := v[0]
	minor := v[1]
	patch := v[2]

	return major, minor, patch
}

func testImgDir() string {
	path, err := filepath.Abs("./cmd/starwars/images")
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func strs2ints(s []string) []int {
	a := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		a[i] = n
	}

	return a
}

func isTest() bool {
	return flag.Lookup("test.v") != nil
}
