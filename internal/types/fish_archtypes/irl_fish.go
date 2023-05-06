package fish_archtypes

import "fmt"

// IRLFish is an interface for fish that exist in the real world
type IRLFish struct {
	Name string
	Type string
}

func (f IRLFish) Details() string {
	return fmt.Sprintf("This is a %v, it is of type %v", f.Name, f.Type)
}
