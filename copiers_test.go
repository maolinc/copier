package copier

import (
	"log"
	"testing"
)

type A struct {
	Num int64
}
type B struct {
	Num string
}

func TestCopiers(t *testing.T) {

	a := A{Num: 41}
	var b B

	_ = Copiers(&b, a)

	log.Println(b)
}
