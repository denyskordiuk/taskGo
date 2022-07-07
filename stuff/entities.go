package stuff

import "fmt"

type Animal struct {
	Name      string
	Weight    int
	FeedCount int
}

type Dog struct {
	Animal
}
type Cat struct {
	Animal
}
type Cow struct {
	Animal
}

type CreateAnimal interface {
	FeedWeight() int
	InfoAnimal() string
}

func (a Animal) FeedWeight() int {
	return a.Weight * a.FeedCount
}

func (a Animal) InfoAnimal() string {
	return fmt.Sprintf("Name - %s\tWeight - %dkg\t Feed - %d\n", a.Name, a.Weight, a.FeedWeight())
}
