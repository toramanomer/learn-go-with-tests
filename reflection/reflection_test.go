package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		name          string
		input         any
		expectedCalls []string
	}{
		{
			name:          "struct with one string fields",
			input:         struct{ Name string }{Name: "Omer"},
			expectedCalls: []string{"Omer"},
		},
		{
			name: "struct with two string fields",
			input: struct {
				Name string
				City string
			}{
				Name: "Omer",
				City: "Istanbul",
			},
			expectedCalls: []string{"Omer", "Istanbul"},
		},
		{
			name: "struct with non string field",
			input: struct {
				Name string
				Age  int
			}{
				Name: "Omer",
				Age:  32,
			},
			expectedCalls: []string{"Omer"},
		},
		{
			name: "nested fields",
			input: Person{
				Name: "Omer",
				Profile: Profile{
					Age:  32,
					City: "Istanbul",
				},
			},
			expectedCalls: []string{"Omer", "Istanbul"},
		},
		{
			name: "pointer",
			input: &Person{
				Name:    "Omer",
				Profile: Profile{Age: 32, City: "Istanbul"},
			},
			expectedCalls: []string{"Omer", "Istanbul"},
		},
		{
			name: "slices",
			input: []Profile{
				{32, "Istanbul"},
				{33, "Bursa"},
			},
			expectedCalls: []string{"Istanbul", "Bursa"},
		},
		{
			name: "arrays",
			input: [2]Profile{
				{32, "Istanbul"},
				{33, "Bursa"},
			},
			expectedCalls: []string{"Istanbul", "Bursa"},
		},
		{
			name: "map",
			input: map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			expectedCalls: []string{"Moo", "Baa"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var got []string

			walk(tt.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.expectedCalls) {
				t.Errorf("got: %v, want: %v", got, tt.expectedCalls)
			}
		})
	}

	t.Run("with channels", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{Age: 32, City: "Istanbul"}
			ch <- Profile{Age: 33, City: "Bursa"}
			close(ch)
		}()

		var got []string
		want := []string{"Istanbul", "Bursa"}

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v, got: %v", want, got)
		}
	})

	t.Run("with function", func(t *testing.T) {

		fn := func() (Profile, Profile) {
			return Profile{32, "Istanbul"}, Profile{33, "Bursa"}
		}

		var got []string
		walk(fn, func(input string) {
			got = append(got, input)
		})

		expected := []string{"Istanbul", "Bursa"}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected: %v, got %v", expected, got)
		}
	})
}
