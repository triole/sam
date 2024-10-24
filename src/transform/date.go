package transform

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	timezone "github.com/gandarez/go-olson-timezone"
)

func (tr Transform) runDate() {
	inputDate, err := tr.strToDate()
	if err != nil {
		logFatal(err, "date processing failure")
	}
	if tr.Conf.Target == "all" {
		printTable(tr.assembleDateTableContent(inputDate))
	} else {
		for _, el := range tr.Impl.DateLayouts {
			if strings.EqualFold(tr.Conf.Target, el.Name) {
				fmt.Printf("%s", inputDate.Format(el.Layout))
			}
		}
	}
}

func (tr Transform) strToDate() (tim time.Time, err error) {
	if tr.Conf.String == "now" {
		tr.Conf.String = tr.now().Format(time.RFC3339Nano)
	}
	if rxMatch("[0-9]{10,}", tr.Conf.String) {
		tim = tr.unixToDate(tr.Conf.String)
		tr.Conf.String = tim.Format(time.RFC3339Nano)
	}
	for _, el := range tr.Impl.DateLayouts {
		tim, err = time.ParseInLocation(
			el.Layout, strings.ToUpper(tr.Conf.String), time.Local,
		)
		if err == nil {
			break
		}
	}
	if tim.Unix() < 0 {
		err = errors.New("can not parse string to date: " + tr.Conf.String)
	}
	return
}

func (tr Transform) getLocation() (zone string, loc *time.Location) {
	var err error
	zone, err = timezone.Name()
	if err != nil {
		log.Fatalf("unable to get timezone name: %v", err)
	}
	loc, err = time.LoadLocation(zone)
	if err != nil {
		logFatal(err, "can not load zone location")
	}
	return
}

func (tr Transform) now() time.Time {
	_, zone := tr.getLocation()
	return time.Now().UTC().In(zone)
}

// func (tr Transform) nowUTC() (now time.Time) {
// 	_, zone := tr.getLocation()
// 	now = time.Now().UTC().In(zone)
// 	return
// }

// func (tr Transform) yesterday() time.Time {
// 	return dp.addDays(-1, dp.today())
// }

// func (tr Transform) today() (today time.Time) {
// 	n := dp.now()
// 	today = time.Date(
// 		n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, dp.TimeZoneLocation,
// 	)
// 	return
// }

// func (tr Transform) tomorrow() time.Time {
// 	return dp.addDays(1, dp.today())
// }

func (tr Transform) assembleDateTableContent(tim time.Time) (r [][]interface{}) {
	header := []interface{}{"Format", "Date"}
	if tr.Conf.Layout {
		header = append(header, "Layout")
	}
	r = append(r, header)
	r = append(r, []interface{}{"UnixTimeStamp", tim.Unix()})
	for _, el := range tr.Impl.DateLayouts {
		if el.Print {
			line := []interface{}{el.Name, tim.Format(el.Layout)}
			if tr.Conf.Layout {
				line = append(line, el.Layout)
			}
			r = append(r, line)
		}
	}
	return
}

func (tr Transform) unixToDate(s string) (tim time.Time) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logFatal(err, "unable to parse string to int: "+s)
	}
	tim = time.Unix(i, 0)
	return
}
