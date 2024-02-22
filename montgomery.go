package montgomery

const (
	P string = "fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f"
)

/*
#cgo LDFLAGS: -lgmp
#include <gmp/gmp.h>

bool hex_to_mpz(const char* hex, mpz_t& m)
{
    uint32_t limbs[8] = { 0 };
    uint32_t count = 256 / 32;
    uint32_t* x = limbs;
    size_t words;
    if (mpz_set_str(m, hex, 16) != 0) {
        return false;
    }
    if (mpz_sizeinbase(m, 2) > count * 32)
    {
        return false;
    }
    mpz_export(x, &words, -1, sizeof(uint32_t), 0, 0, m);
    while (words < count)
    {
        x[words++] = 0;
    }
    return true;
}

uint32_t find_np0(const mpz_t m) {
    uint32_t np0;
    mpz_t temp;
    mpz_init(temp);
    mpz_ui_pow_ui(temp, 2, 32);
    assert(mpz_invert(temp, m, temp));
    np0 = -mpz_get_ui(temp);
    mpz_clear(temp);
    return np0;
}
*/

import "C"

func NP0() uint32 {
	var modulus C.mpz_t
	var np0 C.uint
	C.mpz_init(&modulus[0])
	C.hex_to_mpz((*C.char)(C.CBytes([]byte(P))), &modulus[0])
	np0 = C.find_np0(&modulus[0])
	C.mpz_clear(&modulus[0])
	return np0
}
