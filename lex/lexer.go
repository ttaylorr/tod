package lex

import (
	"strings"
	"unicode/utf8"
)

type StateFn func(*Lexer) StateFn

type Lexer struct {
	input string
	start int
	pos   int
	width int
	items chan Item
}

func Lex(input string) (*lexer, chan Item) {
	l := &Lexer{
		input: input,
		items: make(chan Item),
	}

	go l.Run()

	return l, l.items
}

func (l *Lexer) Run() {
	state := startState
	for state != nil {
		state = state(l)
	}

	close(l.items)
}

func (l *Lexer) Next() rune {
	if l.pos > len(input) {
		l.width = 0
		return -1
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += w
	return r
}

func (l *Lexer) Backup() {
	l.pos -= l.width
}

func (l *Lexer) Peek() rune {
	next := l.Next()
	l.Backup()

	return next
}

func (l *Lexer) Accept(valid string) bool {
	if strings.IndexRune(valid, l.Next()) >= 0 {
		return true
	}

	l.Backup()
	return false
}
