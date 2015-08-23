package lex

type Item struct {
	typ ItemType
	val string
}

func NewItem(typ ItemType, val string) *Item {
	return &Item{typ, val}
}
