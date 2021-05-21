package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"lrg"},
			[]string{"lrg"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"lrg", "shanghai"},
			[]string{"lrg", "shanghai"},
		},
		{
			"Struct with none string field",
			struct {
				Name string
				Age  int
			}{"lrg", 21},
			[]string{"lrg"},
		},
		{
			"new field",
			Person{
				"chj",
				Profile{20, "shanxi"},
			},
			[]string{"chj", "shanxi"},
		}, {
			"Pointers to things",
			&Person{
				"lrg",
				Profile{20, "shanxi"},
			},
			[]string{"lrg", "shanxi"},
		},
		{
			"Slices",
			[]Profile{
				{33, "shanxi"},
				{34, "hunan"},
			},
			[]string{"shanxi", "hunan"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}
