package duck

type QuackBehavior func() string

func QuackLoud() string {
	return "Quack loud"
}

func NoQuack() string {
	return "No quack"
}
