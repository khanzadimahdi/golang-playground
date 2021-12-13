package cmp

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
	data := []Person{
		Person{
			Name: "mahdi",
		},
		Person{
			Name: "ali",
		},
		Person{
			Name: "mohamad",
		},
	}

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	for i, expected := range data {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			result := CreatePerson(expected.Name, expected.Age, expected.DateAdded)
			if diff := cmp.Diff(expected, result, comparer); diff != "" {
				t.Error(diff)
			}
		})
	}
}
