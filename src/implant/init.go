package implant

import (
	"embed"
	_ "embed"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var (
	//go:embed embed
	fs embed.FS
)

type DateLayouts []DateLayout

type DateLayout struct {
	Layout  string
	Matcher string
	Name    string
	Print   bool
}

type Implant struct {
	DateLayouts DateLayouts
}

func Init() (impl Implant) {
	impl.loadLayouts()
	return
}

func (impl *Implant) loadLayouts() {
	var dl DateLayouts
	err := yaml.Unmarshal(readFile("date_layouts.yaml"), &dl)
	if err != nil {
		log.Fatal(err)
	}
	impl.DateLayouts = dl
}

func readFile(path string) (by []byte) {
	by, err := fs.ReadFile(filepath.Join("embed", path))
	if err != nil {
		log.Fatalf("can not load embedded file: %s", err)
	}
	return
}

func (dl DateLayouts) List() (arr []string) {
	for _, el := range dl {
		if el.Print {
			arr = append(arr, el.Name)
		}
	}
	return
}
