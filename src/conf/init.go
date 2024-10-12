package conf

import (
	"reflect"
	"strings"
	"unicode"
)

func Init(cli interface{}) (conf Conf) {
	conf.SubCommand = strings.Split(getcli(cli, "SubCommand", "str").(string), " ")[0]
	cap := capitalize(conf.SubCommand)
	conf.Target = getcli(cli, cap+".Target", "str").(string)
	conf.String = getcli(cli, cap+".Args", "str").(string)
	conf.Length = getcli(cli, cap+".Length", "int").(int)
	conf.Reverse = getcli(cli, cap+".Reverse", "bool").(bool)
	conf.Aggressive = getcli(cli, cap+".Aggressive", "bool").(bool)
	return
}

func New() (conf Conf) {
	return
}

func getcli(cli interface{}, keypath string, expectedFormat string) (r interface{}) {
	key := strings.Split(keypath, ".")
	val := reflect.ValueOf(cli)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if fieldType.Name == key[0] {
			r = field.Interface()
			if len(key) > 1 {
				return getcli(r, key[1], expectedFormat)
			}
			if fieldType.Type.String() == "[]string" {
				arr := field.Interface().([]string)
				r = strings.Join(arr, " ")
			} else {
				r = field.Interface()
			}
		}
	}
	// make sure not to return nil on empty args
	if r == nil {
		switch expectedFormat {
		case "int":
			r = 0
		case "bool":
			r = false
		default:
			r = ""
		}
	}
	return
}

func capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
