package starwars

import (
	"reflect"
	"testing"
)

var CharNames = []string{
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

func Test_charNames(t *testing.T) {
	expect := CharNames
	actual := charNames()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(`expected="%s" actual="%s"`, expect, actual)
	}
}

func Test_randomCharName(t *testing.T) {
	charName := randomCharName()

	isIncluded := false
	for i := 0; i < len(CharNames); i++ {
		if charName == CharNames[i] {
			isIncluded = true
			break
		}
	}

	if !isIncluded {
		t.Errorf("%s is not included in %x", charName, CharNames)
	}
}
