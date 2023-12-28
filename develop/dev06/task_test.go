package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestValidFlagF(t *testing.T) {
	tests := []struct {
		name  string
		input flags
		want  flags
		err   error
	}{
		{"testflagF1", flags{f: "1,5,4,2", s: false, d: " "}, flags{f: "1,5,4,2", s: false, d: " ", fields: []int{1, 2, 4, 5}}, nil},
		{
			"testflagF2", flags{f: "-1,8,1", s: false, d: " "}, flags{f: "-1,8,1", s: false, d: " "},
			errorNegativeNum,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fl = flags{}
				fl = tt.input
				goterr := validFlagF()
				if goterr != tt.err {
					t.Errorf("myCut() goterr = %v , want = %v", goterr, tt.err)
				}
				if !reflect.DeepEqual(fl, tt.want) {
					t.Errorf("флаги не равны")
				}
			},
		)
	}
}

func TestMyCutDelimSpace(t *testing.T) {
	fl = flags{}
	fl.f = "1,3,99"
	fl.d = " "
	fl.fields = []int{1, 3, 99}

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"testcase1", "hello my name is John\ngo is awesome\napple macbook pro", "hello name \ngo awesome \napple pro "},
		{"testcase2", "world world golang\nsafari opera google", "world golang \nsafari google "},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := myCut(bufio.NewScanner(strings.NewReader(tt.input))); got != tt.want {
					t.Errorf("myCut() = %s, want = %s", got, tt.want)
				}
			},
		)
	}
}
