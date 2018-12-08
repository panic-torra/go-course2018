package main

import (
	"fmt"
	"github.com/panic-torra/go-course2018/workshop1/task3/duck"
)

func main() {
	flyingDuck := duck.Duck{FlyBehavior: duck.FlyWithWings{}, QuackBehavior: duck.NoQuack}
	mallocDuck := duck.Duck{FlyBehavior: duck.NoFly{}, QuackBehavior: duck.NoQuack}
	rubberDuck := duck.NewRubberDuck()
	redHeadDuck := duck.NewRedHeadDuck()
	Play(&flyingDuck)
	Play(&mallocDuck)
	Play(rubberDuck)
	Play(redHeadDuck)
}

func Play(d duck.IDuck) {
	fmt.Println(d.Fly())
	fmt.Println(d.Quack())
}
