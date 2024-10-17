package transform

import (
	"sam/src/conf"
	"testing"
)

func TestHash(t *testing.T) {
	s := "hello"
	assertHash(s, "md5", "", 0, "5d41402abc4b2a76b9719d911017c592", t)
	assertHash(s, "blake3", "", 16, "ea8f163db38682925e4491c5e58d4bb3", t)
	assertHash(s, "blake3", "", 32, "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200f", t)
	assertHash(s, "blake3", "", 64, "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338c", t)
	assertHash(s, "blake3", "", 128, "ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338cc68876568ab5c6e524abbcfe881e5b4e1ac9336f3f932d412248c9829536699f07a1b1ce35ffdfe0be5d00c083a8dfd29c9a4303d1374cd70e6abcec6e6b796c", t)
	assertHash(s, "sha1", "", 0, "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", t)
	assertHash(s, "sha256", "", 0, "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824", t)
	assertHash(s, "sha512", "", 0, "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043", t)
	assertHash(s, "whirlpool", "", 0, "0a25f55d7308eca6b9567a7ed3bd1b46327f0f1ffdc804dd8bb5af40e88d78b88df0d002a89e2fdbd5876c523f1b67bc44e9f87047598e7548298ea1c81cfd73", t)

	fil := "../../tests/hash.txt"
	assertHash(s, "md5", fil, 0, "6f5902ac237024bdd0c176cb93063dc4", t)
	assertHash(s, "blake3", fil, 16, "dc5a4edb8240b018124052c330270696", t)
	assertHash(s, "blake3", fil, 32, "dc5a4edb8240b018124052c330270696f96771a63b45250a5c17d3000e823355", t)
	assertHash(s, "sha1", fil, 0, "22596363b3de40b06f981fb85d82312e8c0ed511", t)
	assertHash(s, "sha256", fil, 0, "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447", t)
	assertHash(s, "sha512", fil, 0, "db3974a97f2407b7cae1ae637c0030687a11913274d578492558e39c16c017de84eacdc8c62fe34ee4e12b4b1428817f09b6a2760c3f8a664ceae94d2434a593", t)
	assertHash(s, "whirlpool", fil, 0, "813350ea4f5299958f338ee9c860df4546e7955f18867f5cdd9364960c68e76b15ed599a6d1758a4450d069379ba31fa7193c90ffe2f9d045d42a3526f10aea4", t)

}

func assertHash(str, target, file string, length int, exp string, t *testing.T) {
	conf := conf.New()
	conf.String = str
	conf.Target = target
	conf.File = file
	conf.Length = length
	tr := Init(conf)
	assert(tr.runHash(), exp, t)
}
