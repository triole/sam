package transform

import (
	_ "embed"
	"errors"
	"log"
	"strconv"
	"time"

	timezone "github.com/gandarez/go-olson-timezone"
	yaml "gopkg.in/yaml.v2"
)

var (
	//go:embed embed/date_layouts.yaml
	layouts []byte
)

type dateLayouts []dateLayout

type dateLayout struct {
	Layout  string
	Matcher string
	Name    string
	Print   bool
}

func (tr Transform) runDate() {
	tr.loadLayouts()
	inputDate, err := tr.strToDate()
	if err != nil {
		logFatal(err, "date processing failure")
	}
	printTable(tr.assembleDateTableContent(inputDate))
}

func (tr *Transform) loadLayouts() (dl dateLayouts) {
	err := yaml.Unmarshal(layouts, &dl)
	if err != nil {
		log.Fatal(err)
	}
	tr.DateLayouts = dl
	return
}

func (tr Transform) strToDate() (tim time.Time, err error) {
	if tr.Conf.String == "now" {
		tr.Conf.String = tr.now().Format(time.RFC3339Nano)
	}
	if rxMatch("[0-9]{10,}", tr.Conf.String) {
		tim = tr.unixToDate(tr.Conf.String)
		tr.Conf.String = tim.Format(time.RFC3339Nano)
	}
	for _, el := range tr.DateLayouts {
		tim, err = time.ParseInLocation(
			el.Layout, tr.Conf.String, time.Local,
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
	r = append(r, []interface{}{"Format", "Date", "Layout"})
	r = append(r, []interface{}{"Unix Time Stamp", tim.Unix()})
	for _, el := range tr.DateLayouts {
		if el.Print {
			r = append(r, []interface{}{el.Name, tim.Format(el.Layout), el.Layout})
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
