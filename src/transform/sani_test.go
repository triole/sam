package transform

import (
	"sam/src/conf"
	"sam/src/implant"
	"strings"
	"testing"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func TestSani(t *testing.T) {
	rs := initReplacerSchemes()
	validateSani("HELLO there", "hello_there", rs, t)
	validateSani("hello  there", "hello_there", rs, t)
	validateSani("helläöü", "hellaeoeue", rs, t)
	validateSani("hello's", "hello_s", rs, t)
	validateSani("àáâãå", "aaaaa", rs, t)
	validateSani("èéêë", "eeee", rs, t)
	validateSani("íìîĩï", "iiiii", rs, t)
	validateSani("òóôõ", "oooo", rs, t)
	validateSani("ùúû", "uuu", rs, t)
	validateSani("ç", "c", rs, t)
	validateSani("Hello---World", "hello_world", rs, t)
	validateSani("a---_b", "a_b", rs, t)
	validateSani("a----_----b", "a_b", rs, t)
	validateSani("a-_-b", "a_b", rs, t)
	validateSani("hello--world", "hello-world", rs, t)
	validateSani("hello__world", "hello_world", rs, t)
	validateSani("----", "", rs, t)
	validateSani("a---", "a", rs, t)
	validateSani("a-", "a", rs, t)
	validateSani("hello-__-", "hello", rs, t)
	validateSani("hello____", "hello", rs, t)
	validateSani("hello there   11", "hello_there11", rs, t)
	validateSani("hello there   222", "hello_there_222", rs, t)
	validateSani("hello there   3333", "hello_there_3333", rs, t)
	validateSani("Hello HereAndThere    3333", "hello_here_and_there_3333", rs, t)
}

func TestAccentedCharacters(t *testing.T) {
	rs := initReplacerSchemes()

	// Every character in a group must collapse to exactly one base letter.
	groups := []struct {
		in   string
		want string
	}{
		{"àáâãåāăąǎȁȃȧḁạảǡǻấầẩẫậắằẳẵặ", "a"},
		{"ḃḅḇ", "b"},
		{"çćĉċčḉ", "c"},
		{"ďḋḍḏḑḓđð", "d"},
		{"èéêëēĕėęěȅȇȩḙḛẹẻẽḕḗḝếềểễệ", "e"},
		{"ḟ", "f"},
		{"ĝğġģǧǵḡ", "g"},
		{"ĥȟḣḥḧḩḫẖħ", "h"},
		{"ìíîïĩīĭįǐȉȋḭỉịḯı", "i"},
		{"ĵǰ", "j"},
		{"ķǩḱḳḵ", "k"},
		{"ĺļľḷḻḽḹłŀ", "l"},
		{"ḿṁṃ", "m"},
		{"ñńņňǹṅṇṉṋ", "n"},
		{"òóôõōŏőơǒǫȍȏȯọỏǭȭȱṍṏṑṓốồổỗộớờởỡợ", "o"},
		{"ṕṗ", "p"},
		{"ŕŗřȑȓṙṛṟṝ", "r"},
		{"śŝşšșṡṣṥṧṩẛſ", "s"},
		{"ţťțṫṭṯṱẗŧ", "t"},
		{"ùúûũūŭůűųưǔȕȗṳṵṷụủṹṻứừửữự", "u"},
		{"ṽṿ", "v"},
		{"ŵẁẃẅẇẉẘ", "w"},
		{"ẋẍ", "x"},
		{"ýÿŷȳẏẙỳỵỷỹ", "y"},
		{"źżžẑẓẕ", "z"},
	}
	for _, g := range groups {
		validateSani(g.in, strings.Repeat(g.want, len([]rune(g.in))), rs, t)
	}

	// German umlauts and other Latin letters without canonical decompositions.
	validateSani("äǟ", "aeae", rs, t)
	validateSani("öȫ", "oeoe", rs, t)
	validateSani("üǖǘǚǜ", "ueueueueue", rs, t)
	validateSani("ß", "ss", rs, t)
	validateSani("æǣǽ", "aeaeae", rs, t)
	validateSani("œ", "oe", rs, t)
	validateSani("øǿ", "oo", rs, t)
	validateSani("þ", "th", rs, t)
	validateSani("ÄÖÜẞ", "aeoeuess", rs, t)
	validateSani("ÆØÞ", "aeoth", rs, t)
	validateSani("ĐŁŒ", "dloe", rs, t)

	// Real-world mixed-case names, including uppercase accents.
	validateSani("Mañana", "manana", rs, t)
	validateSani("São Tomé", "sao_tome", rs, t)
	validateSani("Élève à l'hôtel", "eleve_a_l_hotel", rs, t)
	validateSani("Dvořák", "dvorak", rs, t)
	validateSani("Łódź", "lodz", rs, t)
	validateSani("Þingvellir", "thingvellir", rs, t)
	validateSani("Tiếng Việt", "tieng_viet", rs, t)
	validateSani("İstanbul", "istanbul", rs, t)
	validateSani("Œuvre", "oeuvre", rs, t)
	validateSani("Æble Øl Ål", "aeble_ol_al", rs, t)
	validateSani("ÑANDÚ ÁÉÍÓÚ", "nandu_aeiou", rs, t)
	validateSani("Đại Việt", "dai_viet", rs, t)

	// Decomposed (NFD) input leaves combining marks behind after lowercasing.
	validateSani("man\u0303ana", "manana", rs, t)
	validateSani("a\u0301", "a", rs, t)
	validateSani("a\u0308", "ae", rs, t)
	validateSani("o\u0308", "oe", rs, t)
	validateSani("u\u0308", "ue", rs, t)
	validateSani("A\u0308", "ae", rs, t)
}

// Every character whose lowercased NFD form is an ASCII base letter followed
// only by nonspacing marks must normalize back to that base letter. A
// diaeresis on a/o/u is transliterated instead of stripped, matching the
// scheme order.
func TestNFDCoverage(t *testing.T) {
	rs := initReplacerSchemes()
	for r := rune(0); r <= unicode.MaxRune; r++ {
		s := string(r)
		dec := norm.NFD.String(strings.ToLower(s))
		decRunes := []rune(dec)
		if len(decRunes) < 2 || decRunes[0] < 'a' || decRunes[0] > 'z' {
			continue
		}
		allMn := true
		for _, m := range decRunes[1:] {
			if !unicode.Is(unicode.Mn, m) {
				allMn = false
				break
			}
		}
		if !allMn {
			continue
		}
		expected := string(decRunes[0])
		if decRunes[1] == '\u0308' {
			switch decRunes[0] {
			case 'a':
				expected = "ae"
			case 'o':
				expected = "oe"
			case 'u':
				expected = "ue"
			}
		}
		got := sanitize(s, rs)
		if got != expected {
			t.Errorf("U+%04X %q NFD %q: got %q, want %q", r, s, dec, got, expected)
		}
	}
}

func TestSaniTransform(t *testing.T) {
	c := conf.New()
	c.String = "Hello HereAndThere    3333"
	tr := Init(c, implant.Init())
	assert(c, tr.runSani(), "hello_here_and_there_3333", t)
}

func validateSani(inp, exp string, rs tReplacerSchemes, t *testing.T) {
	assert(conf.Conf{String: inp}, sanitize(inp, rs), exp, t)
}
