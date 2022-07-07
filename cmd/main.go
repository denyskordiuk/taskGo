package main

import (
	"fmt"
	"math/rand"
	"taskProject/taskGo/stuff"
	"time"
)

const dogFeed int = 2
const catFeed int = 7
const cowFeed int = 25
const catCoefficient float64 = 0.4
const cowCoefficient float64 = 2.45

func farmGenerator(animalCount int) (animals []stuff.CreateAnimal) {
	var animal stuff.CreateAnimal

	for i := 0; i < animalCount; i++ {
		animalNumber := rand.Intn(3) + 1
		animalWeightTemplate := rand.Intn(20) + 40

		switch animalNumber {
		case 1:
			animal = stuff.Dog{stuff.Animal{
				Name:      "Dog",
				Weight:    animalWeightTemplate,
				FeedCount: dogFeed,
			}}
		case 2:
			animal = stuff.Cat{stuff.Animal{
				Name:      "Cat",
				Weight:    int(float64(animalWeightTemplate) * catCoefficient),
				FeedCount: catFeed,
			}}
		case 3:
			animal = stuff.Cow{stuff.Animal{
				Name:      "Cow",
				Weight:    int(float64(animalWeightTemplate) * cowCoefficient),
				FeedCount: cowFeed,
			}}
		}
		animals = append(animals, animal)
	}

	return animals
}

func farmInfo(animals []stuff.CreateAnimal) {
	fmt.Println("\nSo here is animals list:")
	for i, a := range animals {
		fmt.Printf("#%d\t%s", i+1, a.InfoAnimal())
	}
}

func allFeedCount(animals []stuff.CreateAnimal) (count int) {
	for _, a := range animals {
		count += a.FeedWeight()
	}
	return count
}

func main() {
	rand.Seed(time.Now().UnixNano())
	animalNumber := 30
	animals := farmGenerator(animalNumber)
	farmInfo(animals)
	fmt.Printf("\nCount of feed u need for this animals - %d\n", allFeedCount(animals))

}
