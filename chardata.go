package main

import (
	"fmt"
)

type charData struct {
	acct string
	char string
	game string
}

func (cd charData) toString() string {
	return fmt.Sprintf("%v,%v,%v\n", cd.char, cd.acct, cd.game)
}
