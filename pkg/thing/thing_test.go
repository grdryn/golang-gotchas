package thing

import (
	"reflect"
	"testing"
)

func TestThing(t *testing.T) {
	scenarios := []struct {
		name   string
		given  Thing
		expect Thing
		mutate func(given Thing) (Mutated Thing)
	}{
		{
			name:   "when no mutation, should be same object",
			given:  MakeThing(),
			expect: MakeThing(),
			mutate: func(given Thing) Thing { return given },
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			actual := scenario.mutate(scenario.given)
			if !reflect.DeepEqual(actual, scenario.expect) {
				t.Fatalf(
					"Actual vs. expected differs:\nActual: %v\nExpected: %v",
					actual,
					scenario.expect,
				)
			}
		})
	}
}
