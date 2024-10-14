package conf

import (
	"reflect"
	"strings"
	"unicode"
)

func Init(cli interface{}) (conf Conf) {
	conf.SubCommand = strings.Split(getcli(cli, "SubCommand", "str").(string), " ")[0]
	cap := capitalize(conf.SubCommand)
	conf.String = getcli(cli, cap+".Args", "").(string)
	conf.Target = getcli(cli, cap+".Target", "").(string)
	conf.SubString = getcli(cli, cap+".Substring", "").(string)
	conf.Language = getcli(cli, cap+".Language", "english").(string)
	conf.Length = getcli(cli, cap+".Length", 0).(int)
	conf.Reverse = getcli(cli, cap+".Reverse", false).(bool)
	conf.Precision = getcli(cli, cap+".Precision", "").(int)
	conf.Aggressive = getcli(cli, cap+".Aggressive", false).(bool)
	return
}

func New() (conf Conf) {
	return
}

func getcli(cli interface{}, keypath string, defaultReturnValue interface{}) (r interface{}) {
	key := strings.Split(keypath, ".")
	val := reflect.ValueOf(cli)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if fieldType.Name == key[0] {
			r = field.Interface()
			if len(key) > 1 {
				return getcli(r, key[1], defaultReturnValue)
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
		r = defaultReturnValue
	}
	return
}

func capitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
