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
	mod, ok := new(big.Int).SetString(P, 16)
	require.True(t, ok)
	base := big.NewInt(math.MaxInt64)
	exp := big.NewInt(10000)
	result, err := powm_odd(base, exp, mod)
	require.Nil(t, err)

	// verify result
	correct := new(big.Int).Set(base)
	correct.Exp(correct, exp, mod)
	require.Equal(t, result.Cmp(correct), 0)

	// print result
	fmt.Println("powm", result.Text(16))
}
