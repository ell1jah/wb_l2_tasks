package main

import (
	"reflect"
	"testing"
)

func TestUnpackingStr(t *testing.T) {
	tests := []struct {
		name   string
		strarg string
		want   string
		err    error
	}{
		{"test1", "a4bc2d5e", "aaaabccddddde", nil},
		{"test2", "abcd", "abcd", nil},
		{"test3", "45", "", ErrorIncorrectString},
		{"test4", "", "", nil},
		{"test5", `qwe\4\5`, "qwe45", nil},
		{"test6", `qwe\45`, "qwe44444", nil},
		{"test7", `qwe\\5`, `qwe\\\\\`, nil},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				got, goterr := UnpackingStr(test.strarg)
				if !reflect.DeepEqual(goterr, test.err) {
					t.Errorf("UnpackingStr() goterr = %v , want = %v", goterr, test.err)
				}
				if got != test.want {
					t.Errorf("UnpackingStr() got = %v , want = %v", got, test.want)
				}
			},
		)
	}
}
