package transform

import (
	_ "embed"
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
	yaml "gopkg.in/yaml.v2"
)

//go:embed sani.yaml
var replacerSchemesStr string

// replacerSchemes are compiled once, they are applied in the order they are
// defined in sani.yaml.
var replacerSchemes = initReplacerSchemes()

type tReplacerSchemes []tReplacerScheme

type tReplacerScheme struct {
	Rx *regexp.Regexp // Precompiled regex pattern
	Nu string         // Replacement string
}

// tReplacerSchemeConfig is used for YAML unmarshaling
type tReplacerSchemeConfig struct {
	Rx string `yaml:"rx"`
	Nu string `yaml:"nu"`
}

var (
	saniAcronymRx = regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	saniCamelRx   = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

func (tr Transform) runSani() (r string) {
	return sanitize(tr.Conf.String, replacerSchemes)
}

// sanitize replaces odd characters by their closest simple alphabetical
// representation and normalizes the remaining separators. The result only
// contains lowercase ascii letters, digits, dashes and underscores.
func sanitize(str string, replacerSchemes tReplacerSchemes) (r string) {
	r = splitCamelCase(str)
	r = strings.ToLower(r)
	// Decompose accents so the scheme's \p{Mn} rule can strip the marks.
	r = norm.NFD.String(r)
	for _, rs := range replacerSchemes {
		r = rs.Rx.ReplaceAllString(r, rs.Nu)
	}
	return
}

// splitCamelCase inserts underscores at camel case word boundaries, so
// "HereAndThere" becomes "Here_And_There". Other separators are left to the
// replacer schemes.
func splitCamelCase(str string) string {
	str = saniAcronymRx.ReplaceAllString(str, "${1}_${2}")
	return saniCamelRx.ReplaceAllString(str, "${1}_${2}")
}

func initReplacerSchemes() (r tReplacerSchemes) {
	var configSchemes []tReplacerSchemeConfig
	err := yaml.Unmarshal([]byte(replacerSchemesStr), &configSchemes)
	if err != nil {
		panic("sani: failed to parse replacer schemes: " + err.Error())
	}
	r = make(tReplacerSchemes, len(configSchemes))
	for i, cfg := range configSchemes {
		rx, err := regexp.Compile(cfg.Rx)
		if err != nil {
			panic("sani: invalid replacer scheme: " + err.Error())
		}
		r[i] = tReplacerScheme{
			Rx: rx,
			Nu: cfg.Nu,
		}
	}
	return
}
