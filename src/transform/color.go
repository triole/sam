package transform

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type tColorCodes struct {
	Hex   string
	RGB   []uint8
	RGBA  []uint8
	CMYK  []uint8
	YCbCr []uint8
}

func (tr Transform) PrintColorInfo(s string) string {
	cc := tr.colorInfo(s)
	return tr.printableColorCodesString(cc)
}

func (tr Transform) colorInfo(s string) (cc tColorCodes) {
	var err error
	args := strings.Split(s, " ")
	if len(args) == 1 {
		cc.Hex = args[0]
		if !strings.HasPrefix(cc.Hex, "#") {
			cc.Hex = "#" + cc.Hex
		}
		cc.RGBA, err = tr.hexToRGBA(cc.Hex)
		if err != nil {
			fmt.Printf("%s: %+v\n", "can not convert hex to rgba", err)
		}
	}
	if len(args) >= 3 {
		cc.RGBA, _ = tr.makeRGBA(args)
		cc.Hex = tr.rgbToHex(cc.RGBA[0], cc.RGBA[1], cc.RGBA[2])
	}
	cc.RGB = cc.RGBA[:len(cc.RGBA)-1]
	c, m, y, k := color.RGBToCMYK(cc.RGB[0], cc.RGB[1], cc.RGB[2])
	cc.CMYK = []uint8{c, m, y, k}
	y, cb, cr := color.RGBToYCbCr(cc.RGB[0], cc.RGB[1], cc.RGB[2])
	cc.CMYK = []uint8{c, m, y, k}
	cc.YCbCr = []uint8{y, cb, cr}
	return
}

func (tr Transform) makeRGBA(arr []string) (r []uint8, err error) {
	var i int
	var uintArr []uint8
	for _, el := range arr {
		i, err = strconv.Atoi(el)
		if err == nil {
			uintArr = append(uintArr, uint8(i))
		} else {
			err = errors.New("invalid rgba string")
		}
	}
	r = []uint8{uintArr[0], uintArr[1], uintArr[2], 255}
	return r, err
}

func (tr Transform) hexToRGBA(s string) (r []uint8, err error) {
	var c color.RGBA
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = errors.New("invalid length, must be 7 or 4")
	}
	r = []uint8{c.R, c.G, c.B, c.A}
	return
}

func (tr Transform) rgbToHex(r, g, b uint8) string {
	return fmt.Sprintf("0x%02x%02x%02x", r, g, b)
}

func (tr Transform) printableColorCodesString(cc tColorCodes) (r string) {
	r = fmt.Sprintf("%6s %s\n", "Hex", cc.Hex)
	r += fmt.Sprintf("%6s %+v\n", "RGB", cc.RGB)
	r += fmt.Sprintf("%6s %+v\n", "RGBA", cc.RGBA)
	r += fmt.Sprintf("%6s %+v\n", "CMYK", cc.CMYK)
	r += fmt.Sprintf("%6s %+v", "YCbCr", cc.YCbCr)
	return
}
