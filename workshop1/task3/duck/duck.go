package duck

type IDuck interface {
	Fly() string
	Quack() string
}

type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

func (d *Duck) Fly() string {
	return d.FlyBehavior.FlyImpl()
}

func (d Duck) Quack() string {
	return d.QuackBehavior()
}

func NewRubberDuck() IDuck {
	return &Duck{FlyBehavior: FlyWithWings{}, QuackBehavior: QuackLoud}
}

type RedHeadDuck struct {
	Duck
}

func NewRedHeadDuck() *RedHeadDuck {
	duck := RedHeadDuck{Duck{FlyBehavior: FlyWithWings{}, QuackBehavior: QuackLoud}}
	duck.QuackBehavior = func() string {
		return "I'm red head duck!"
	}

	return &duck
}
