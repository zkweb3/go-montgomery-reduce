package montgomery

import (
	"fmt"
	"testing"
)

func TestGetNp0(t *testing.T) {
	np0 := NP0()
	fmt.Println("np0", np0)
}
