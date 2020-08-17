package model

import "time"

type Board struct {
	ID       uint
	Name     string
	Position uint
}

type Section struct {
	ID       uint
	Name     string
	Position uint
	Board    *Board
	Cards    []*Card
}

type Card struct {
	ID          uint
	Name        string
	Position    uint
	Description string
	Section     *Section
	Labels      []*Label
	LimitTime   *time.Time
	Images      []string
}

type Label struct {
	ID      uint
	Name    string
	BoardID uint
}
