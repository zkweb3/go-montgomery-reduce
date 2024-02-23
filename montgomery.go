package montgomery

/*
#cgo CFLAGS: -I/usr/local/gmp/include
#cgo LDFLAGS: -L/usr/local/gmp/lib -lgmp
#include <stdint.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include <gmp.h>

int hex_to_mpz(const char* hex, mpz_t* m)
{
    uint32_t limbs[8] = { 0 };
    uint32_t count = 256 / 32;
    uint32_t* x = limbs;
    size_t words;
    if (mpz_set_str(*m, hex, 16) != 0) {
        return -1;
    }
    if (mpz_sizeinbase(*m, 2) > count * 32)
    {
        return -1;
    }
    mpz_export(x, &words, -1, sizeof(uint32_t), 0, 0, *m);
    while (words < count)
    {
        x[words++] = 0;
    }
    return 0;
}

int mpz_to_hex(const mpz_t r, char** hex) {
    int size = (r->_mp_size * sizeof(uint64_t)) * 2;
    *hex = (char*)malloc(size + 1);
    memcpy(*hex, mpz_get_str(NULL, 16, r), size);
    return size;
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
import (
	"math/big"
	"unsafe"
)

func NP0(m *big.Int) uint32 {
	var modulus C.mpz_t
	var np0 C.uint
	C.mpz_init(&modulus[0])
	C.hex_to_mpz((*C.char)(C.CBytes([]byte(m.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&modulus[0])))
	np0 = C.find_np0(&modulus[0])
	C.mpz_clear(&modulus[0])
	return uint32(np0)
}

func powm_odd(base, exp, mod *big.Int) (*big.Int, error) {
	var rop C.mpz_t
	var b, e, m C.mpz_t
	var r *C.char
	var n C.int
	C.mpz_init(&rop[0])
	C.hex_to_mpz((*C.char)(C.CBytes([]byte(base.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&b[0])))
	C.hex_to_mpz((*C.char)(C.CBytes([]byte(exp.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&e[0])))
	C.hex_to_mpz((*C.char)(C.CBytes([]byte(mod.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&m[0])))
	C.mpz_powm(&rop[0], &b[0], &e[0], &m[0])
	n = C.mpz_to_hex(&rop[0], &r)
	br := C.GoBytes(unsafe.Pointer(r), n)
	result := new(big.Int).SetBytes(br)
	C.free(unsafe.Pointer(r))
	C.mpz_clear(&rop[0])
	return result, nil
}
