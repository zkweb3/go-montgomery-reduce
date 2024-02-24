package montgomery

/*
#cgo CFLAGS: -I/usr/local/gmp/include
#cgo LDFLAGS: -L/usr/local/gmp/lib -lgmp
#include <stdint.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>
#include <stdio.h>
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
    strcpy(*hex, mpz_get_str(NULL, 16, r));
    return strlen(*hex);
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

uint32_t bn2mont(mpz_t mont, const mpz_t bn, const mpz_t n) {
    mpz_t    temp;
    uint32_t n0, inv;
    if (mpz_tstbit(n, 0) == 0) {
        return 0;
    }
    if (mpz_cmp(bn, n) >= 0) {
        return 0;
    }
    mpz_init(temp);
    mpz_mul_2exp(temp, bn, 8 * 32);
    mpz_fdiv_r(mont, temp, n);
    mpz_clear(temp);

    n0 = mpz_get_ui(n);
    inv = n0 * (n0 * n0 + 14);
    inv = inv * (inv * n0 + 2);
    inv = inv * (inv * n0 + 2);
    inv = inv * (inv * n0 + 2);
    return inv;
}

void mont2bn(mpz_t bn, const mpz_t mont, const mpz_t n, uint32_t np0) {
    mpz_t    add, p0, p1, p2;
    int32_t  index;
    uint32_t low;

    puts(mpz_get_str(NULL, 16, mont));
    fflush(stdout);
    puts(mpz_get_str(NULL, 16, n));
    fflush(stdout);
    printf("%u\n", np0);
    fflush(stdout);
    if (np0 * (uint32_t)mpz_get_ui(n) != 0xFFFFFFFF) {
        return;
    }
    mpz_inits(add, p0, p1, p2, NULL);
    mpz_set(p0, mont);
    for (index = 0; index < 8; index++) {
        low = np0 * (uint32_t)mpz_get_ui(p0);
        mpz_mul_ui(add, n, low);
        mpz_add(p1, p0, add);
        mpz_fdiv_q_2exp(p2, p1, 32);
        mpz_swap(p0, p2);
        mpz_clear(p2);
        mpz_init(p2);
    }
    puts(mpz_get_str(NULL, 16, p0));
    fflush(stdout);
    if (mpz_cmp(p0, n) < 0)
        mpz_set(bn, p0);
    else
        mpz_sub(bn, p0, n);
    mpz_clears(add, p0, p1, p2, NULL);
}
*/
import "C"
import (
    "errors"
    "fmt"
    "math/big"
    "unsafe"
)

func init_mpz(zs ...C.mpz_t) {
    for _, z := range zs {
        C.mpz_init(&z[0])
    }
}

func clear_mpz(zs ...C.mpz_t) {
    for _, z := range zs {
        C.mpz_clear(&z[0])
    }
}

func NP0(m *big.Int) uint32 {
    var modulus C.mpz_t
    var np0 C.uint

    C.mpz_init(&modulus[0])
    defer C.mpz_clear(&modulus[0])
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(m.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&modulus[0])))
    np0 = C.find_np0(&modulus[0])
    return uint32(np0)
}

func powm_odd(base, exp, mod *big.Int) (*big.Int, error) {
    var rop, b, e, m C.mpz_t
    var ptr *C.char
    var len C.int

    init_mpz(rop, b, e, m)
    defer clear_mpz(rop, b, e, m)
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(base.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&b[0])))
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(exp.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&e[0])))
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(mod.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&m[0])))
    C.mpz_powm(&rop[0], &b[0], &e[0], &m[0])
    len = C.mpz_to_hex(&rop[0], &ptr)
    defer C.free(unsafe.Pointer(ptr))
    fmt.Println("len", int(len))
    br := C.GoBytes(unsafe.Pointer(ptr), len)
    r, ok := new(big.Int).SetString(*(*string)(unsafe.Pointer(&br)), 16)
    if !ok {
        return nil, errors.New("convert error")
    }
    return r, nil
}

func bn2mont(bn, mod *big.Int) (*big.Int, uint32) {
    var mont, b, m C.mpz_t
    var np0 C.uint
    var ptr *C.char
    var len C.int

    init_mpz(mont, b, m)
    defer clear_mpz(mont, b, m)
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(bn.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&b[0])))
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(mod.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&m[0])))
    np0 = C.bn2mont(&mont[0], &b[0], &m[0])
    len = C.mpz_to_hex(&mont[0], &ptr)
    defer C.free(unsafe.Pointer(ptr))
    fmt.Println("np0", int(np0))
    fmt.Println("len", int(len))
    br := C.GoBytes(unsafe.Pointer(ptr), len)
    r, ok := new(big.Int).SetString(*(*string)(unsafe.Pointer(&br)), 16)
    if !ok {
        return nil, 0
    }
    return r, uint32(np0)
}

func mont2bn(mont, mod *big.Int, np0 uint32) (*big.Int, error) {
    var bn, mt, m C.mpz_t
    var ptr *C.char
    var len C.int

    init_mpz(bn, mt, m)
    defer clear_mpz(bn, mt, m)
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(mont.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&mt[0])))
    C.hex_to_mpz((*C.char)(C.CBytes([]byte(mod.Text(16)))), (*C.mpz_t)(unsafe.Pointer(&m[0])))
    C.mont2bn(&bn[0], &mt[0], &m[0], C.uint(np0))
    len = C.mpz_to_hex(&bn[0], &ptr)
    defer C.free(unsafe.Pointer(ptr))
    fmt.Println("len", int(len))
    br := C.GoBytes(unsafe.Pointer(ptr), len)
    r, ok := new(big.Int).SetString(*(*string)(unsafe.Pointer(&br)), 16)
    if !ok {
        return nil, errors.New("convert error")
    }
    return r, nil
}
