package model

type Board struct {
	ID   uint
	Name string
}

type Section struct {
	ID    uint
	Name  string
	Board *Board
	Cards []*Card
}

type Card struct {
	ID          uint
	Name        string
	Description string
	Section     *Section
}
