package oh_my_go

type mammal interface {
	canWalk() bool
	canFeed() bool
}

func newMan() mammal {
	return man{}
}

func newDog() mammal {
	return dog{}
}
