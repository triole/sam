package transform

import (
	_ "embed"
	"errors"
	"log"
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
	tr.DateLayouts = tr.loadLayouts()
	inputDate := tr.strToDate()
	printTable(tr.assembleDateTableContent(inputDate))
}

func (tr Transform) loadLayouts() (dl dateLayouts) {
	err := yaml.Unmarshal(layouts, &dl)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (tr Transform) detectLayout(str string) (r dateLayout, err error) {
	for _, el := range tr.DateLayouts {
		if rxMatch(el.Matcher, str) {
			r = el
			break
		}
	}
	if r.Layout == "" {
		err = errors.New("no fitting date layout for: " + str)
	}
	return
}

func (tr Transform) strToDate() (tim time.Time) {
	if tr.Conf.String == "now" {
		tr.Conf.String = tr.now().Format(time.RFC3339Nano)
	}
	layout, err := tr.detectLayout(tr.Conf.String)
	if err != nil {
		logFatal(err, "detect layout failed")
	} else {
		tim, err = time.ParseInLocation(
			layout.Layout, tr.Conf.String, time.Local,
		)
		if err != nil {
			logFatal(err, "unable to parse string to date, layout name: "+layout.Name)
		}
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
