package calculator

import (
	. "github.com/ovechkin-dm/mockio/mock"
	"testing"
)

func TestSimple_Mockio(t *testing.T) {
	SetUp(t)
	m := Mock[Logger]()
	m.Log(10)
	Verify(m, AtLeastOnce()).Log(10)
	// Verify(m, Once()).Log(Equal(10))
}
