package conf

import (
	"reflect"
	"strings"
	"unicode"
)

func Init(cli interface{}) (conf Conf) {
	conf.SubCommand = strings.Split(getcli(cli, "SubCommand").(string), " ")[0]
	cap := capitalize(conf.SubCommand)
	conf.Target = getcli(cli, cap+".Target").(string)
	conf.String = getcli(cli, cap+".Args").(string)
	conf.Length = getcli(cli, cap+".Length").(int)
	conf.Reverse = getcli(cli, cap+".Reverse").(bool)
	conf.Aggressive = getcli(cli, cap+".Aggressive").(bool)
	return
}

func InitTest() (conf Conf) {
	return
}

func getcli(cli interface{}, keypath string) (r interface{}) {
	key := strings.Split(keypath, ".")
	val := reflect.ValueOf(cli)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if fieldType.Name == key[0] {
			r = field.Interface()
			if len(key) > 1 {
				return getcli(r, key[1])
			}
			if fieldType.Type.String() == "[]string" {
				arr := field.Interface().([]string)
				r = strings.Join(arr, " ")
			} else {
				r = field.Interface()
			}
		}
	}
	// make sure not to return nil on empty bools
	if r == nil {
		r = false
	}
	return
}

func capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
