package transform

import (
	"fmt"
	"strings"
)

func (tr Transform) runMatch() (r string) {
	hash := strings.Split(tr.Conf.String, " ")[0]
	tr.Conf.String = strings.Join(strings.Split(tr.Conf.String, " ")[1:], " ")
	switch tr.Conf.Target {
	case "md5":
		r = tr.md5Match(hash)
	case "sha1":
		r = tr.sha1Match(hash)
	case "sha256":
		r = tr.sha256Match(hash)
	case "sha384":
		r = tr.sha384Match(hash)
	case "sha512":
		r = tr.sha512Match(hash)
	case "blake3":
		r = tr.blake3Match(hash)
	case "whirlpool":
		r = tr.whirlpoolMatch(hash)
	case "rake":
		r = tr.rakeMatch(hash)
	}
	return
}

func (tr Transform) md5Match(hash string) string {
	compHash := tr.md5()
	return fmt.Sprintf("%s :: %s :: %v", tr.Conf.String, hash, compHash == hash)
}

func (tr Transform) sha1Match(hash string) string {
	compHash := tr.sha1()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) sha256Match(hash string) string {
	compHash := tr.sha256()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) sha384Match(hash string) string {
	compHash := tr.sha384()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) sha512Match(hash string) string {
	compHash := tr.sha512()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) blake3Match(hash string) string {
	compHash := tr.blake3()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) whirlpoolMatch(hash string) string {
	compHash := tr.whirlpool()
	return fmt.Sprintf("%v", compHash == hash)
}

func (tr Transform) rakeMatch(hash string) string {
	compHash := tr.rake()
	return fmt.Sprintf("%v", compHash == hash)
}
