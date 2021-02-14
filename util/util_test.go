package util

import (
	"strings"
	"testing"
)

func TestGetVersions(t *testing.T) {
	major, minor, patch := GetVersions("1.2.3")
	if !(major == 1 && minor == 2 && patch == 3) {
		t.Error(major, minor, patch)
	}

	major, minor, patch = GetVersions("0.2.3")
	if !(major == 0 && minor == 2 && patch == 3) {
		t.Error(major, minor, patch)
	}

	major, minor, patch = GetVersions("8.9.0")
	if !(major == 8 && minor == 9 && patch == 0) {
		t.Error(major, minor, patch)
	}

	major, minor, patch = GetVersions("10.0.3")
	if !(major == 10 && minor == 0 && patch == 3) {
		t.Error(major, minor, patch)
	}
}

func TestImgFilePath(t *testing.T) {
	expect := "cmd/starwars/images/bb-8.png"
	actual := ImgFilePath("bb-8")

	if !strings.HasSuffix(actual, expect) {
		t.Errorf(`"%s" does not have suffix "%s"`, actual, expect)
	}
}
