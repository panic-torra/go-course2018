package duck

type FlyBehavior interface {
	FlyImpl() string
}

type FlyWithWings struct{}

func (f FlyWithWings) FlyImpl() string {
	return "Fly with wings"
}

type NoFly struct{}

func (f NoFly) FlyImpl() string {
	return "No fly"
}
