package tymbol

import (
	"fmt"
	"math"
	"strings"
)

const (
	SPACE    = " "
	NEW_LINE = "\n"
)

type Table struct {
	Title   string
	headers []string
	columns [][]string
	Options Options

	cellLength   int
	tableLength  int
	maxColLength []int
	maxRowLength []int

	canvas strings.Builder
}

func NewTable(title string, headers []string, columns [][]interface{}) (Table, error) {
	if len(columns) == 0 {
		return Table{}, fmt.Errorf("Columns cannot be empty!")
	}

	if len(headers) > 0 && len(headers) != len(columns) {
		return Table{}, fmt.Errorf("Number of headers and columns don't match: %d %d", len(headers), len(columns))
	}

	var colLength int
	for i := range columns {
		if i == 0 {
			colLength = len(columns[i])
			continue
		}

		if colLength != len(columns[i]) {
			return Table{}, fmt.Errorf("Columns must be same lenght. Assumed len: %d. Diff len column index: %d", colLength, i)
		}
	}
	numberOfRows := len(columns[0])
	if len(headers) > 0 {
		numberOfRows += 1
	}

	maxColLength := make([]int, len(columns))
	maxRowLength := make([]int, numberOfRows)

	for i := range headers {
		maxColLength[i] = len(headers[i])
		if len(headers) > 0 && len(headers[i]) > maxRowLength[0] {
			maxRowLength[0] = len(headers[i])
		}
	}

	strColumns := make([][]string, len(columns))
	for i := range columns {
		for j := range columns[i] {
			val := fmt.Sprintf("%v", columns[i][j])
			if maxColLength[i] < len(val) {
				maxColLength[i] = len(val)
			}
			var extraIndex int
			if len(headers) > 0 {
				extraIndex = 1
			}
			if maxRowLength[j+extraIndex] < len(val) {
				maxRowLength[j+extraIndex] = len(val)
			}
			strColumns[i] = append(strColumns[i], val)
		}
	}

	t := Table{
		Title:        title,
		headers:      headers,
		columns:      strColumns,
		Options:      defaultOptions(),
		maxColLength: maxColLength,
		maxRowLength: maxRowLength,
	}

	return t, nil
}

func (t *Table) ResetCanvas() {
	t.canvas.Reset()
}

func (t *Table) Draw() string {

	if !t.Options.CellFitContent() {
		t.cellLength = 2*t.Options.CellPadding() + t.Options.CellLength()
		t.tableLength = t.cellLength*len(t.columns) + len(t.columns) + 1
	} else {
		for i := 0; i < len(t.columns); i++ {
			t.tableLength += 2*t.Options.CellPadding() + t.maxColLength[i]
		}
		t.tableLength += len(t.columns) + 1
	}

	t.drawTitle()
	t.drawHeader()
	t.drawBody()
	return t.canvas.String()
}

func (t *Table) newLine() {
	t.canvas.WriteString(NEW_LINE)
}

func (t *Table) getLengthByIndex(idx int) int {
	if t.Options.CellFitContent() {
		return t.maxColLength[idx] + 2*t.Options.CellPadding()
	}
	return t.cellLength
}

func (t *Table) drawLine(hasLeft, hasRight bool, crossSym, hSym rune, cellLength int) {
	if hasLeft {
		t.canvas.WriteRune(crossSym)
	}

	for i := 0; i < cellLength; i++ {
		t.canvas.WriteRune(hSym)
	}
	if hasRight {
		t.canvas.WriteRune(crossSym)
	}
}

func (t *Table) drawValueLine(hasLeft, hasRight bool, vSym rune, lineAlign align, cellLength int, v string) {
	if hasLeft {
		t.canvas.WriteRune(vSym)
	}

	var left, right int
	switch lineAlign {
	case CENTER:
		left = (cellLength - len(v)) / 2
		right = cellLength - len(v) - left
	case LEFT:
		left = t.Options.CellPadding()
		right = cellLength - left - len(v)
	case RIGHT:
		left = cellLength - t.Options.CellPadding() - len(v)
		right = t.Options.CellPadding()
	}

	for i := 0; i < left; i++ {
		t.canvas.WriteString(SPACE)
	}
	t.canvas.WriteString(v)
	for i := 0; i < right; i++ {
		t.canvas.WriteString(SPACE)
	}
	if hasRight {
		t.canvas.WriteRune(vSym)
	}
}

func (t *Table) drawValueMultiLine(hasLeft, hasRight bool, vSym rune, lineAlign align, cellLength int, rowHeight int, filledRows int, cursor int, v string) {
	upperPadding := (rowHeight - filledRows) / 2
	if cursor < upperPadding || cursor >= filledRows+upperPadding {
		t.drawValueLine(hasLeft, hasRight, vSym, lineAlign, cellLength, SPACE)
	} else {
		position := (cursor - upperPadding) * t.Options.CellLength()
		charsLeft := len(v) - position
		if charsLeft < t.Options.CellLength() {
			t.drawValueLine(hasLeft, hasRight, vSym, lineAlign, cellLength, v[position:])
		} else {
			t.drawValueLine(hasLeft, hasRight, vSym, lineAlign, cellLength, v[position:position+t.Options.CellLength()])
		}
	}

}

func (t *Table) drawHeader() {
	if len(t.headers) == 0 {
		return
	}

	for i := 0; i < len(t.headers); i++ {
		hasLeft, hasRight := false, true
		if i == 0 {
			hasLeft, hasRight = true, true
		}
		t.drawLine(hasLeft, hasRight, t.Options.CrossHeaderSym(), t.Options.HHeaderSym(), t.getLengthByIndex(i))
	}
	t.newLine()
	rowHeight := int(math.Ceil(float64(t.maxRowLength[0]) / float64(t.Options.CellLength())))
	for j := 0; j < rowHeight; j++ {
		for i := 0; i < len(t.headers); i++ {
			hasLeft, hasRight := false, true
			if i == 0 {
				hasLeft, hasRight = true, true
			}
			if !t.Options.CellFitContent() {
				filledRows := int(math.Ceil(float64(len(t.headers[i])) / float64(t.Options.CellLength())))
				t.drawValueMultiLine(hasLeft, hasRight, t.Options.VHeaderSym(), t.Options.HeaderAlign(), t.getLengthByIndex(i), rowHeight, filledRows, j, t.headers[i])
			} else {
				t.drawValueLine(hasLeft, hasRight, t.Options.VHeaderSym(), t.Options.HeaderAlign(), t.getLengthByIndex(i), t.headers[i])
			}
		}
		t.newLine()
	}
	for i := 0; i < len(t.headers); i++ {
		hasLeft, hasRight := false, true
		if i == 0 {
			hasLeft, hasRight = true, true
		}
		t.drawLine(hasLeft, hasRight, t.Options.CrossHeaderSym(), t.Options.HHeaderSym(), t.getLengthByIndex(i))
	}
	t.newLine()
}

func (t *Table) drawBody() {
	headerIndex := 1
	if len(t.headers) == 0 {
		headerIndex = 0
		for i := range t.columns {
			hasLeft, hasRight := false, true
			if i == 0 {
				hasLeft, hasRight = true, true
			}
			t.drawLine(hasLeft, hasRight, t.Options.CrossLineSym(), t.Options.HLineSym(), t.getLengthByIndex(i))
		}
		t.newLine()
	}

	for i := 0; i < len(t.columns[0]); i++ {
		rowHeight := int(math.Ceil(float64(t.maxRowLength[i+headerIndex]) / float64(t.Options.CellLength())))
		for n := 0; n < rowHeight; n++ {
			for j := range t.columns {
				hasLeft, hasRight := false, true
				if j == 0 {
					hasLeft, hasRight = true, true
				}
				if !t.Options.CellFitContent() {
					filledRows := int(math.Ceil(float64(len(t.columns[j][i])) / float64(t.Options.CellLength())))
					t.drawValueMultiLine(hasLeft, hasRight, t.Options.VLineSym(), t.Options.CellAlign(), t.getLengthByIndex(j), rowHeight, filledRows, n, t.columns[j][i])
				} else {
					t.drawValueLine(hasLeft, hasRight, t.Options.VLineSym(), t.Options.CellAlign(), t.getLengthByIndex(j), t.columns[j][i])
				}
			}
			t.newLine()
		}
		for j := range t.columns {
			hasLeft, hasRight := false, true
			if j == 0 {
				hasLeft, hasRight = true, true
			}
			t.drawLine(hasLeft, hasRight, t.Options.CrossLineSym(), t.Options.HLineSym(), t.getLengthByIndex(j))
		}
		t.newLine()
	}
}

func (t *Table) drawTitle() {
	if t.Title == "" {
		return
	}

	var left, right int
	switch t.Options.TitleAlign() {
	case CENTER:
		left = (t.tableLength - len(t.Title)) / 2
		right = t.tableLength - len(t.Title) - left
	case LEFT:
		left = 0
		right = t.tableLength - len(t.Title)
	case RIGHT:
		left = t.tableLength - len(t.Title)
		right = 0
	}

	for i := 0; i < left; i++ {
		t.canvas.WriteString(SPACE)
	}

	t.canvas.WriteString(t.Title)

	for i := 0; i < right; i++ {
		t.canvas.WriteString(SPACE)
	}
	t.newLine()
}
