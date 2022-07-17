package stuff

import (
	"fmt"
	"math/rand"
)

const (
	catFeed      int = 7
	minCatWeight int = 3
)

type Cat struct {
	name       string
	weight     int
	isFavorite bool
}

func NewCat(name string) *Cat {
	weight := minCatWeight + rand.Intn(10)

	return &Cat{
		name:       "cat-" + name,
		weight:     weight,
		isFavorite: weight > 10,
	}
}

func (c *Cat) FeedWeight() int {
	if c.isFavorite {
		return c.weight * catFeed * 2
	}

	return c.weight * catFeed
}

func (c *Cat) Info() string {
	return fmt.Sprintf("Name: %s\t Weight: %d\t Favorite: %t\t Feed:%d", c.name, c.weight, c.isFavorite, c.FeedWeight())
}
