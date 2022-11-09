package transform

import (
	"log"
	"strconv"
	"strings"
)

func separateFirstArg(s string) (string, string) {
	arr := strings.Split(s, " ")
	return arr[0], strings.Join(arr[1:], " ")
}

func logFatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func toInt(itf interface{}) (i int) {
	switch val := itf.(type) {
	case string:
		v, err := strconv.Atoi(val)
		logFatal(err, "Can not convert "+val)
		i = v
	case int:
		i = val
	}
	return
}
