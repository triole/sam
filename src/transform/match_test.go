package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"strconv"
	"strings"
	"testing"
)

func TestMatch(t *testing.T) {
	s := "hello"
	assertMatch(
		"5d41402abc4b2a76b9719d911017c592"+" "+s, "md5", 0, t,
	)
	assertMatch(
		"ea8f163db38682925e4491c5e58d4bb3"+" "+s, "blake3", 16, t,
	)
	assertMatch(
		"ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200f"+" "+s, "blake3", 32, t,
	)
	assertMatch(
		"ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338c"+" "+s, "blake3", 64, t,
	)
	assertMatch(
		"ea8f163db38682925e4491c5e58d4bb3506ef8c14eb78a86e908c5624a67200fe992405f0d785b599a2e3387f6d34d01faccfeb22fb697ef3fd53541241a338cc68876568ab5c6e524abbcfe881e5b4e1ac9336f3f932d412248c9829536699f07a1b1ce35ffdfe0be5d00c083a8dfd29c9a4303d1374cd70e6abcec6e6b796c"+" "+s, "blake3", 128, t,
	)
	assertMatch(
		"aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"+" "+s, "sha1", 0, t,
	)
	assertMatch(
		"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"+" "+s, "sha256", 0, t,
	)
	assertMatch(
		"9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"+" "+s, "sha512", 0, t,
	)
	assertMatch(
		"0a25f55d7308eca6b9567a7ed3bd1b46327f0f1ffdc804dd8bb5af40e88d78b88df0d002a89e2fdbd5876c523f1b67bc44e9f87047598e7548298ea1c81cfd73"+" "+s, "whirlpool", 0, t,
	)
	assertMatch(
		"obinaP2ouXcuwDRNdgu1157eDZ5dqDeyqDTZ"+" "+s, "rake", 8, t,
	)
	assertMatch(
		"obinaP2ouXcuwDRNdgu1157eDZ5dqDeyqDTZejqqPJ27flueaXqrs4rqoFBZ218RJR1FDV"+" "+s, "rake", 16, t,
	)
}

func assertMatch(inpStr, target string, length int, t *testing.T) {
	runAssert(inpStr, target, length, true, t)
	inpStr = strings.Replace(inpStr, "a", "b", -1)
	inpStr = strings.Replace(inpStr, "b", "c", -1)
	inpStr = strings.Replace(inpStr, "c", "d", -1)
	inpStr = strings.Replace(inpStr, "d", "e", -1)
	inpStr = strings.Replace(inpStr, "e", "f", -1)
	runAssert(inpStr, target, length, false, t)
}

func runAssert(inpStr, target string, length int, exp bool, t *testing.T) {
	hash := strings.Split(inpStr, " ")[0]
	str := strings.Join(strings.Split(inpStr, " ")[1:], " ")
	conf := conf.New()
	conf.String = inpStr
	conf.Target = target
	conf.Length = length
	expStr := str + " :: " + hash + " :: " + strconv.FormatBool(exp)
	tr := Init(conf, implant.Init())
	assert(conf, tr.runMatch(), expStr, t)
}
