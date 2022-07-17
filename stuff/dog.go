package stuff

import (
	"fmt"
	"math/rand"
)

const (
	dogFeed      int = 2
	minDogWeight int = 10
)

type Dog struct {
	name   string
	weight int
}

func NewDog(name string) *Dog {
	return &Dog{
		name:   "dog-" + name,
		weight: minDogWeight + rand.Intn(20),
	}
}

func (d *Dog) FeedWeight() int {
	return d.weight * dogFeed
}

func (d *Dog) Info() string {
	return fmt.Sprintf("Name: %s\t Weight: %d\t Feed:%d", d.name, d.weight, d.FeedWeight())
}
