// все случаи want проверены с помощью утилиты sort с соответсвующим флагом
package main

import (
	"reflect"
	"testing"
)

func Test_Flag_k_mySort(t *testing.T) {
	fl.k = 2
	fl.n = false
	fl.r = false
	fl.u = false
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"test_1 flag k", []string{"dog 6", "cat 1", "snake 3"}, []string{"cat 1", "snake 3", "dog 6"}},
		{"test_2 flag k", []string{"99ddd 7 g", "6bb 5 g", "aaac 1 h"}, []string{"aaac 1 h", "6bb 5 g", "99ddd 7 g"}},
		{"test_3 flag k", []string{"44 h", "tiger v", "apple 3"}, []string{"apple 3", "44 h", "tiger v"}},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := mySort(tt.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mySort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_Flag_n_mySort(t *testing.T) {
	fl.k = 0
	fl.n = true
	fl.r = false
	fl.u = false
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"test_1 flag n", []string{"1dog 6", "44cat 1", "7snake 3"}, []string{"1dog 6", "7snake 3", "44cat 1"}},
		{"test_2 flag n", []string{"88", "974", "22"}, []string{"22", "88", "974"}},
		{"test_3 flag n", []string{"ggwp", "1pink", "jj"}, []string{"ggwp", "jj", "1pink"}},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := mySort(tt.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mySort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_Flag_r_mySort(t *testing.T) {
	fl.k = 0
	fl.n = false
	fl.r = true
	fl.u = false
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"test_1 flag r", []string{"1dog 6", "44cat 1", "7snake 3"}, []string{"7snake 3", "44cat 1", "1dog 6"}},
		{"test_2 flag r", []string{"88", "974", "22"}, []string{"974", "88", "22"}},
		{"test_3 flag r", []string{"ggwp", "1pink", "78jj"}, []string{"ggwp", "78jj", "1pink"}},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := mySort(tt.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mySort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_Flag_u_mySort(t *testing.T) {
	fl.k = 0
	fl.n = false
	fl.r = false
	fl.u = true
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			"test_1 flag u", []string{"44cat 1", "1dog 6", "1dog 6", "7snake 3"}, []string{
				"1dog 6", "44cat 1",
				"7snake 3",
			},
		},
		{"test_2 flag u", []string{"88", "974", "22", "22"}, []string{"22", "88", "974"}},
		{"test_3 flag u", []string{"ggwp", "1pink", "1pink", "1pink", "78jj"}, []string{"1pink", "78jj", "ggwp"}},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := mySort(tt.input); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("mySort() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
