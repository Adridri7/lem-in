package src

type Population struct {
	InitialPop      int
	CurrentPop      int
	AntsCoordinates []Position
}

type Position struct {
	X, Y int
}
