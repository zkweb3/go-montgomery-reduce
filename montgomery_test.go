package montgomery

import (
    "fmt"
    "math"
    "math/big"
    "testing"

    "github.com/stretchr/testify/require"
)

const (
    P  string = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f"
    bn string = "b0007aa30c1c50f8"
)

func TestMontNp0(t *testing.T) {
    mod, ok := new(big.Int).SetString(P, 16)
    require.True(t, ok)
    gn, ok := new(big.Int).SetString(bn, 16)
    require.True(t, ok)

    mont, np0 := bn2mont(gn, mod)
    require.NotEqual(t, np0, 0)
    fmt.Println("np0", np0)

    require.Equal(t, NP0(mod), np0)
    fmt.Println("mont", mont.Text(16))

    ggn, err := mont2bn(mont, mod, np0)
    require.Nil(t, err)
    require.Equal(t, ggn.Cmp(gn), 0)
    fmt.Println("bn", ggn.Text(16))
}

func TestPowMod(t *testing.T) {
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
