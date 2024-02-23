package montgomery

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	P string = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f"
)

func TestGetNp0(t *testing.T) {
	m, ok := new(big.Int).SetString(P, 16)
	require.True(t, ok)
	np0 := NP0(m)
	fmt.Println("np0", np0)
}

func TestPowm(t *testing.T) {
	m, ok := new(big.Int).SetString(P, 16)
	require.True(t, ok)
	b := big.NewInt(math.MaxInt64)
	e := big.NewInt(10000)
	p, _ := powm_odd(b, e, m)
	fmt.Println("powm", p.Text(16))
}
