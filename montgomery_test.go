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
    bx string = "b0007aa30c1c50f8"
    by string = "30ffffffefffffbb"
)

func TestMontMul(t *testing.T) {
    mod, ok := new(big.Int).SetString(P, 16)
    require.True(t, ok)
    bigx, ok := new(big.Int).SetString(bx, 16)
    require.True(t, ok)
    bigy, ok := new(big.Int).SetString(by, 16)
    require.True(t, ok)

    montx, np0 := bn2mont(bigx, mod)
    require.NotEqual(t, np0, 0)
    fmt.Println("np0", np0)
    monty, _ := bn2mont(bigy, mod)
    require.Equal(t, NP0(mod), np0)
    fmt.Println("mont_x", montx.Text(16))
    fmt.Println("mont_y", monty.Text(16))

    montz, err := mont_mul(montx, monty, mod, np0)
    require.Nil(t, err)
    fmt.Println("mont_z", montz.Text(16))
    bigz, err := mont2bn(montz, mod, np0)
    require.Nil(t, err)
    fmt.Println("z", bigz.Text(16))

    // verify result
    correct := new(big.Int).Mul(bigx, bigy)
    correct.Mod(correct, mod)
    require.Equal(t, correct.Cmp(bigz), 0)
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

func BenchmarkMontMul(b *testing.B) {
    mod, ok := new(big.Int).SetString(P, 16)
    require.True(b, ok)
    bigx, ok := new(big.Int).SetString(bx, 16)
    require.True(b, ok)
    bigy, ok := new(big.Int).SetString(by, 16)
    require.True(b, ok)

    montx, np0 := bn2mont(bigx, mod)
    require.NotEqual(b, np0, 0)
    fmt.Println("np0", np0)
    monty, _ := bn2mont(bigy, mod)
    require.Equal(b, NP0(mod), np0)
    fmt.Println("mont_x", montx.Text(16))
    fmt.Println("mont_y", monty.Text(16))

    // b.StartTimer()
	// for i := 0; i < b.N; i++ {
    //     _, err := mont_mul(montx, monty, mod, np0)
    //     require.Nil(b, err)
	// }
	// b.StopTimer()
}
