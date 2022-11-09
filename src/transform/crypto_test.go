package transform

import "testing"

var (
	v = "hello"
)

func TestMd5(t *testing.T) {
	assert(tr.Md5(v), "5d41402abc4b2a76b9719d911017c592", t)
}

func TestBlake3(t *testing.T) {
	assert(tr.Blake3("16 "+v), "ea8f163db38682925e4491c5e58d4bb3", t)
	assert(tr.Blake3("32 "+v), "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200f", t)
	assert(tr.Blake3("64 "+v), "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338c", t)
	assert(tr.Blake3("128 "+v), "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338cc68876568ab5c6e524abbcfe881e5b4e1ac9336f3f932d412248c9829536699f07a1b1ce35ffdfe0be5d00c083a8dfd29c9a4303d1374cd70e6abcec6e6b796c", t)
}

func TestSha1(t *testing.T) {
	assert(tr.Sha1(v), "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", t)
}

func TestSha256(t *testing.T) {
	assert(tr.Sha256(v), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", t)
}

func TestSha512(t *testing.T) {
	assert(tr.Sha512(v), "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043", t)
}

func TestRipemd160(t *testing.T) {
	assert(tr.Ripemd160(v), "108f07b8382412612c048d07d13f814118445acd", t)
}

func TestWhirlpool(t *testing.T) {
	assert(tr.Whirlpool(v), "0a25f55d7308eca6b9567a7ed3bd1b46327f0f1ffdc804dd8bb5af40e88d78b88df0d002a89e2fdbd5876c523f1b67bc44e9f87047598e7548298ea1c81cfd73", t)
}
