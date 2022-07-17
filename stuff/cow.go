package stuff

import (
	"fmt"
	"math/rand"
)

const (
	cowFeed        int = 25
	reducedCowFeed int = 20
	minCowWeight   int = 175
)

type Cow struct {
	name   string
	weight int
}

func NewCow(name string) *Cow {
	return &Cow{
		name:   "cow-" + name,
		weight: minCowWeight + rand.Intn(50),
	}
}

func (c *Cow) FeedWeight() int {
	if c.weight >= 200 {
		return c.weight * reducedCowFeed
	}

	return c.weight * cowFeed
}

func (c *Cow) Info() string {
	return fmt.Sprintf("Name: %s\t Weight: %d\t Feed:%d", c.name, c.weight, c.FeedWeight())
}
