package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/denyskordiuk/taskGo/stuff"
	nameGenerator "github.com/goombaio/namegenerator"
)

var (
	errEmptyFarm = errors.New("empty farm")
)

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	generator := nameGenerator.NewNameGenerator(seed)
	animals := farmGenerator(10, generator)
	for _, animal := range animals {
		fmt.Println(animal.Info())
	}

	count, err := allFeedCount(animals)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCount of feed u need for this animals - %d\n", count)
}

func farmGenerator(count int, g nameGenerator.Generator) []stuff.Animal {
	animals := make([]stuff.Animal, 0, count)
	for i := 0; i < count; i++ {
		animalNumber := rand.Intn(3)
		name := g.Generate()

		switch animalNumber {
		case 0:
			animals = append(animals, stuff.NewCow(name))
		case 1:
			animals = append(animals, stuff.NewDog(name))
		case 2:
			animals = append(animals, stuff.NewCat(name))
		}
	}

	return animals
}

func allFeedCount(animals []stuff.Animal) (int, error) {
	if len(animals) == 0 {
		return 0, errEmptyFarm
	}

	var count int

	for _, animal := range animals {
		count += animal.FeedWeight()
	}

	return count, nil
}
