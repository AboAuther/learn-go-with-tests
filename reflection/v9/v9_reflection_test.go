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
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		}, {
			"Arrays",
			[2]Profile{
				{20, "shanxi"},
				{21, "ningxia"},
			},
			[]string{"shanxi", "ningxia"},
		},
	}
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"shanghai": "MaritimeUniversity",
			"beijing":  "Electric",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "MaritimeUniversity")
		assertContains(t, got, "Electric")
	})
	t.Run("with Channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{20, "shanxi"}
			aChannel <- Profile{21, "shanghai"}
			close(aChannel)
		}()
		var got []string
		want := []string{"shanxi", "shanghai"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
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

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("excepted %+v to contain %q but it didn't", haystack, needle)
	}
}
