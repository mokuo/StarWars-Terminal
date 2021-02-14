package util

import (
	"reflect"
	"testing"
)

var ExpectCharNames = []string{
	"bb-8",
	"boba-fett",
	"c-3po",
	"chewbacca",
	"darth-vader",
	"darth_maul",
	"princess_amidala",
	"r2-d2",
	"royal_guard",
	"stormtrooper",
	"the_emperor",
	"yoda",
}

func TestRandomCharName(t *testing.T) {
	charName := RandomCharName()

	isIncluded := false
	for i := 0; i < len(ExpectCharNames); i++ {
		if charName == ExpectCharNames[i] {
			isIncluded = true
			break
		}
	}

	if !isIncluded {
		t.Errorf("%s is not included in %x", charName, ExpectCharNames)
	}
}

func TestCharNames(t *testing.T) {
	expect := ExpectCharNames
	actual := CharNames()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`expected="%s" actual="%s"`, expect, actual)
	}
}
