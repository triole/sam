package transform

import (
	_ "embed"
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
	Desc    string
	Layout  string
	Matcher string
}

func (tr Transform) runDate() {
	input := tr.strToDate()
	printTable(tr.assembleTableContent(input))
}

func (tr Transform) loadLayouts() (dl dateLayouts) {
	err := yaml.Unmarshal(layouts, &dl)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (tr Transform) detectLayout(str string) (r string) {
	dl := tr.loadLayouts()
	for _, el := range dl {
		if rxMatch(el.Matcher, str) {
			r = el.Layout
			break
		}
	}
	return
}

func (tr Transform) strToDate() (tim time.Time) {
	if tr.Conf.String == "now" {
		tr.Conf.String = tr.now().Format(time.UnixDate)
	}
	layout := tr.detectLayout(tr.Conf.String)
	if layout == "UnixDate" {
		layout = time.UnixDate
	}
	var err error
	tim, err = time.ParseInLocation(
		layout, tr.Conf.String, time.Local,
	)
	if err != nil {
		log.Fatalf("can not parse string to date: %v", err)
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

func (tr Transform) assembleTableContent(tim time.Time) (r [][]interface{}) {
	r = append(r, []interface{}{"Format", "Date"})
	r = append(r, []interface{}{"Unix Time Stamp", tim.Unix()})
	r = append(r, []interface{}{"Unix Date", tim.Format(time.UnixDate)})
	r = append(r, []interface{}{"RFC3339", tim.Format(time.RFC3339)})
	r = append(r, []interface{}{"RFC822Z", tim.Format(time.RFC822Z)})
	r = append(r, []interface{}{"RFC1123Z", tim.Format(time.RFC1123Z)})
	return
}
