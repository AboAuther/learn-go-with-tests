package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExceptedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"lrg"},
			[]string{"lrg"},
		}, {
			"struct with two string field",
			struct {
				Name string
				city string
			}{"lrg", "shanghai"},
			[]string{"lrg", "shanghai"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExceptedCalls) {
				t.Errorf("got %v, want %v", got, test.ExceptedCalls)
			}
		})
	}
}
