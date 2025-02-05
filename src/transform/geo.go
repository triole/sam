package transform

import (
	"fmt"
	"strings"

	"github.com/pierrre/geohash"
)

func (tr Transform) runGeo() (r string) {
	switch tr.Conf.Target {
	case "enc":
		r = tr.geoEncode()
	case "dec":
		tr.geoDecode()
	}
	return
}

func (tr Transform) geoEncode() string {
	arr := strings.Split(tr.Conf.String, " ")
	var fl1, fl2 float64
	var err error
	if len(arr) > 1 {
		fl1, err = stringToFloat(arr[0])
		if err != nil {
			fmt.Printf("%+v\n", "can not convert string to float: "+arr[0])
		}
		fl2, err = stringToFloat(arr[1])
		if err != nil {
			fmt.Printf("%+v\n", "can not convert string to float: "+arr[1])
		}
	}
	return geohash.Encode(fl1, fl2, tr.Conf.Length)
}

func (tr Transform) geoDecode() {
	h, err := geohash.Decode(tr.Conf.String)
	if err == nil {
		t := tr.geoAssembleDateTableContent(h)
		printTable(t)
	}
}

func (tr Transform) geoAssembleDateTableContent(ghb geohash.Box) (r [][]interface{}) {
	header := []interface{}{"", "Cordinates"}
	if tr.Conf.Layout {
		header = append(header, "Layout")
	}
	r = append(r, header)
	r = append(r, []interface{}{"Location", ghb.Round()})
	r = append(r, []interface{}{"Center", ghb.Center()})
	r = append(r, []interface{}{"Latitude", ghb.Lat})
	r = append(r, []interface{}{"Longitude", ghb.Lon})
	return
}
