package main

import (
	"reflect"
	"testing"
)

func TestMySearchAnagram(t *testing.T) {
	teststr := []string{"пятак", "листок", "слиток", "тяпка", "столик", "пятка", "propper", ""}
	want := &map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}
	got := MySearchAnagram(teststr)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mySearchAnagram() got = %v , want = %v", got, want)
	}
}
