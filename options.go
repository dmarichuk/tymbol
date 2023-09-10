package table

import "fmt"

const (
	LEFT   = "left"
	RIGHT  = "right"
	CENTER = "center"
)

var availableAligns = [3]string{LEFT, RIGHT, CENTER}

type align = string

type Options struct {
	titleAlign align

	cellFitContent bool
	cellLength     int
	cellPadding    int
	cellAlign      align
	crossLineSym   rune
	vLineSym       rune
	hLineSym       rune

	headerAlign    align
	crossHeaderSym rune
	hHeaderSym     rune
	vHeaderSym     rune
}

func defaultOptions() Options {
	return Options{
		titleAlign:     "center",
		cellFitContent: false,
		cellLength:     10,
		cellPadding:    2,
		headerAlign:    "center",
		crossHeaderSym: '#',
		hHeaderSym:     '=',
		vHeaderSym:     '#',
		cellAlign:      "center",
		crossLineSym:   '+',
		hLineSym:       '-',
		vLineSym:       '|',
	}
}

func checkAlignOption(a string) bool {
	for i := 0; i < len(availableAligns); i++ {
		if a == availableAligns[i] {
			return true
		}
	}
	return false
}

func (o *Options) TitleAlign() string {
	return o.titleAlign
}

func (o *Options) SetTitleAlign(a align) error {
	if ok := checkAlignOption(a); !ok {
		return fmt.Errorf("Unknown align option. Expected %v, got %s", availableAligns, a)
	}
	o.titleAlign = a
	return nil
}

func (o *Options) HeaderAlign() string {
	return o.headerAlign
}

func (o *Options) SetHeaderAlign(a align) error {
	if ok := checkAlignOption(a); !ok {
		return fmt.Errorf("Unknown align option. Expected %v, got %s", availableAligns, a)
	}
	o.headerAlign = a
	return nil
}

func (o *Options) CellLength() int {
	return o.cellLength
}

func (o *Options) SetCellLength(l int) error {
	if l <= 0 {
		return fmt.Errorf("Value must be greater than 0")
	}
	o.cellLength = l
	return nil
}

func (o *Options) CellFitContent() bool {
	return o.cellFitContent
}

func (o *Options) SetCellFitContent(p bool) error {
	o.cellFitContent = p
	return nil
}

func (o *Options) CellPadding() int {
	return o.cellPadding
}

func (o *Options) SetCellPadding(p int) error {
	if p < 0 {
		return fmt.Errorf("Value must be positive")
	}
	o.cellPadding = p
	return nil
}

func (o *Options) CellAlign() align {
	return o.cellAlign
}

func (o *Options) SetCellAlign(a align) error {
	if ok := checkAlignOption(a); !ok {
		return fmt.Errorf("Unknown align option. Expected %v, got %s", availableAligns, a)
	}
	o.cellAlign = a
	return nil
}

func (o *Options) CrossHeaderSym() rune {
	return o.crossHeaderSym
}

func (o *Options) SetCrossHeaderSym(s rune) error {
	o.crossHeaderSym = s
	return nil
}

func (o *Options) HHeaderSym() rune {
	return o.hHeaderSym
}

func (o *Options) SetHHeaderSym(s rune) error {
	o.hHeaderSym = s
	return nil
}
func (o *Options) VHeaderSym() rune {
	return o.vHeaderSym
}

func (o *Options) SetVHeaderSym(s rune) error {
	o.vHeaderSym = s
	return nil
}

func (o *Options) CrossLineSym() rune {
	return o.crossLineSym
}

func (o *Options) SetCrossLineSym(s rune) error {
	o.crossLineSym = s
	return nil
}

func (o *Options) VLineSym() rune {
	return o.vLineSym
}

func (o *Options) SetVLineSym(s rune) error {
	o.vLineSym = s
	return nil
}

func (o *Options) HLineSym() rune {
	return o.hLineSym
}

func (o *Options) SetHLineSym(s rune) error {
	o.hLineSym = s
	return nil
}
