package thing

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
