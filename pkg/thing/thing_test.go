package thing

// Copyright (C) 2020 Gerard Ryan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
		{
			name:  "when labels modified, subThing labels should not be implicitly modified",
			given: MakeThing(),
			expect: Thing{
				Labels: map[string]string{"key": "value"},
				SubThing: SubThing{
					Labels: map[string]string{},
				},
			},
			mutate: func(given Thing) Thing {
				given.Labels["key"] = "value"
				// given.SubThing.Labels should be empty...
				return given
			},
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
