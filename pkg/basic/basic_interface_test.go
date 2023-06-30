package basic

import (
	"fmt"
	"testing"
)

func Test_BasicInterface(t *testing.T) {
	x := AnimalQuark(&dog{})
	x.quark()
}

func Test_(t *testing.T) {
	fmt.Println("Hello Test")
}
