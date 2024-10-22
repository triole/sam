package transform

import (
	_ "embed"
	"fmt"
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

func (tr Transform) runDate() (r string) {
	input := tr.strToDate()
	return tr.printableDateStrings(input)
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
		log.Fatal("can not load zone location: %v", err)
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

func (tr Transform) printableDateStrings(tim time.Time) (r string) {
	r = fmt.Sprintf("%14s %d\n", "UnixTimeStamp", tim.Unix())
	r += fmt.Sprintf("%14s %s\n", "UnixDate", tim.Format(time.UnixDate))
	r += fmt.Sprintf("%14s %s\n", "RFC822", tim.Format(time.RFC822))
	r += fmt.Sprintf("%14s %s\n", "RFC822Z", tim.Format(time.RFC822Z))
	r += fmt.Sprintf("%14s %s\n", "RFC1123", tim.Format(time.RFC1123))
	r += fmt.Sprintf("%14s %s\n", "RFC1123Z", tim.Format(time.RFC1123Z))
	r += fmt.Sprintf("%14s %s\n", "RFC3339", tim.Format(time.RFC3339))
	return
}
