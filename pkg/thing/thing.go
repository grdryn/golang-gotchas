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

import "fmt"

// Thing has labels and a SubThing
type Thing struct {
	Labels   map[string]string
	SubThing SubThing
}

func (t Thing) String() string {
	return fmt.Sprintf("Thing{\n\tLabels: %v,\n\tSubThing: %v,\n}", t.Labels, t.SubThing)
}

// SubThing is like Thing, but just has labels
type SubThing struct {
	Labels map[string]string
}

func (s SubThing) String() string {
	return fmt.Sprintf("SubThing{ Labels: %v }", s.Labels)
}

// MakeThing makes a Thing
func MakeThing() Thing {
	defaultLabels := map[string]string{}

	return Thing{
		Labels: defaultLabels,
		SubThing: SubThing{
			Labels: defaultLabels,
		},
	}
}
