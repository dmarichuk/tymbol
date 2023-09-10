package tymbol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {

	t.Run("Successful creation", func(t *testing.T) {
		tab, _ := NewTable(
			"test",
			[]string{"h1", "h2", "h3"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
		)

		assert.Equal(t, tab, Table{
			Title:        "test",
			headers:      []string{"h1", "h2", "h3"},
			columns:      [][]string{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
			maxColLength: []int{3, 3, 3},
			maxRowLength: []int{2, 3, 3, 3},
			Options: Options{
				titleAlign:     "center",
				headerAlign:    "center",
				crossHeaderSym: '#',
				vHeaderSym:     '#',
				hHeaderSym:     '=',
				cellLength:     10,
				cellPadding:    2,
				cellAlign:      "center",
				crossLineSym:   '+',
				vLineSym:       '|',
				hLineSym:       '-',
			},
		})
	})

	t.Run("Header and Columns doesn't match", func(t *testing.T) {
		_, err := NewTable(
			"test",
			[]string{"h1", "h2"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
		)
		if assert.Error(t, err) {
			assert.Equal(t, "Number of headers and columns don't match: 2 3", err.Error())
		}
	})

	t.Run("Columns are different size", func(t *testing.T) {
		_, err := NewTable(
			"test",
			[]string{"h1", "h2"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22"}},
		)
		if assert.Error(t, err) {
			assert.Equal(t, "Columns must be same lenght. Assumed len: 3. Diff len column index: 1", err.Error())
		}
	})

	t.Run("Columns are empty", func(t *testing.T) {
		_, err := NewTable(
			"test",
			[]string{"h1", "h2"},
			[][]interface{}{},
		)
		if assert.Error(t, err) {
			assert.Equal(t, "Columns cannot be empty!", err.Error())
		}
	})

}

func TestDrawTable(t *testing.T) {
	t.Run("Default table", func(t *testing.T) {

		tab, _ := NewTable(
			"",
			[]string{"h1", "h2", "h3"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
		)

		got := tab.Draw()

		want := `#==============#==============#==============#
#      h1      #      h2      #      h3      #
#==============#==============#==============#
|     c11      |     c21      |     c31      |
+--------------+--------------+--------------+
|     c12      |     c22      |     c32      |
+--------------+--------------+--------------+
|     c13      |     c23      |     c33      |
+--------------+--------------+--------------+
`
		assert.Equal(t, want, got)
	})

	t.Run("Headerless table", func(t *testing.T) {

		tab, _ := NewTable(
			"",
			[]string{},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
		)
		got := tab.Draw()

		want := `+--------------+--------------+--------------+
|     c11      |     c21      |     c31      |
+--------------+--------------+--------------+
|     c12      |     c22      |     c32      |
+--------------+--------------+--------------+
|     c13      |     c23      |     c33      |
+--------------+--------------+--------------+
`
		assert.Equal(t, want, got)
	})

}

func TestTableChangeSymbols(t *testing.T) {
	tab, _ := NewTable(
		"",
		[]string{"h1", "h2", "h3"},
		[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
	)
	tab.Options.SetHHeaderSym('o')
	tab.Options.SetVHeaderSym('o')
	tab.Options.SetCrossHeaderSym('o')
	tab.Options.SetVLineSym('[')
	tab.Options.SetHLineSym('-')
	tab.Options.SetCrossLineSym('-')

	expectedTable := Table{
		Title:        "",
		headers:      []string{"h1", "h2", "h3"},
		columns:      [][]string{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}, {"c31", "c32", "c33"}},
		maxColLength: []int{3, 3, 3},
		maxRowLength: []int{2, 3, 3, 3},
		Options: Options{
			titleAlign:     "center",
			headerAlign:    "center",
			cellLength:     10,
			cellPadding:    2,
			cellAlign:      "center",
			crossHeaderSym: 'o',
			vHeaderSym:     'o',
			hHeaderSym:     'o',
			crossLineSym:   '-',
			vLineSym:       '[',
			hLineSym:       '-',
		},
	}
	assert.Equal(t, expectedTable, tab)

	got := tab.Draw()
	want := `oooooooooooooooooooooooooooooooooooooooooooooo
o      h1      o      h2      o      h3      o
oooooooooooooooooooooooooooooooooooooooooooooo
[     c11      [     c21      [     c31      [
----------------------------------------------
[     c12      [     c22      [     c32      [
----------------------------------------------
[     c13      [     c23      [     c33      [
----------------------------------------------
`
	assert.Equal(t, want, got)
}

func TestTableAlign(t *testing.T) {

	t.Run("LEFT align", func(t *testing.T) {
		tab, _ := NewTable(
			"test",
			[]string{"h1", "h2"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}},
		)

		err := tab.Options.SetHeaderAlign(LEFT)
		assert.Equal(t, nil, err)
		err = tab.Options.SetCellAlign(LEFT)
		assert.Equal(t, nil, err)
		err = tab.Options.SetTitleAlign(LEFT)
		assert.Equal(t, nil, err)

		expectedTable := Table{
			Title:        "test",
			headers:      []string{"h1", "h2"},
			columns:      [][]string{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}},
			maxColLength: []int{3, 3},
			maxRowLength: []int{2, 3, 3, 3},
			Options: Options{
				titleAlign:     "left",
				headerAlign:    "left",
				cellLength:     10,
				cellPadding:    2,
				cellAlign:      "left",
				crossHeaderSym: '#',
				vHeaderSym:     '#',
				hHeaderSym:     '=',
				crossLineSym:   '+',
				vLineSym:       '|',
				hLineSym:       '-',
			},
		}
		assert.Equal(t, expectedTable, tab)

		got := tab.Draw()
		want := `test                           
#==============#==============#
#  h1          #  h2          #
#==============#==============#
|  c11         |  c21         |
+--------------+--------------+
|  c12         |  c22         |
+--------------+--------------+
|  c13         |  c23         |
+--------------+--------------+
`
		assert.Equal(t, want, got)
	})

	t.Run("RIGHT align", func(t *testing.T) {
		tab, _ := NewTable(
			"test",
			[]string{"h1", "h2"},
			[][]interface{}{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}},
		)
		err := tab.Options.SetCellAlign(RIGHT)
		assert.Equal(t, nil, err)
		err = tab.Options.SetTitleAlign(RIGHT)
		assert.Equal(t, nil, err)
		err = tab.Options.SetHeaderAlign(RIGHT)
		assert.Equal(t, nil, err)

		expectedTable := Table{
			Title:        "test",
			headers:      []string{"h1", "h2"},
			columns:      [][]string{{"c11", "c12", "c13"}, {"c21", "c22", "c23"}},
			maxColLength: []int{3, 3},
			maxRowLength: []int{2, 3, 3, 3},
			Options: Options{
				titleAlign:     "right",
				cellLength:     10,
				cellPadding:    2,
				headerAlign:    "right",
				cellAlign:      "right",
				crossHeaderSym: '#',
				vHeaderSym:     '#',
				hHeaderSym:     '=',
				crossLineSym:   '+',
				vLineSym:       '|',
				hLineSym:       '-',
			},
		}
		assert.Equal(t, expectedTable, tab)

		got := tab.Draw()
		want := `                           test
#==============#==============#
#          h1  #          h2  #
#==============#==============#
|         c11  |         c21  |
+--------------+--------------+
|         c12  |         c22  |
+--------------+--------------+
|         c13  |         c23  |
+--------------+--------------+
`
		assert.Equal(t, want, got)
	})
}

func TestFitContentTable(t *testing.T) {
	t.Run("Fit content table", func(t *testing.T) {

		tab, _ := NewTable(
			"",
			[]string{"id", "key", "value"},
			[][]interface{}{{1, 2, 3}, {"test1", "test2", "test3"}, {3.14, 0.3, 1.0000003}},
		)
		tab.Options.SetCellFitContent(true)
		got := tab.Draw()

		want := `#======#=========#=============#
#  id  #   key   #    value    #
#======#=========#=============#
|  1   |  test1  |    3.14     |
+------+---------+-------------+
|  2   |  test2  |     0.3     |
+------+---------+-------------+
|  3   |  test3  |  1.0000003  |
+------+---------+-------------+
`
		assert.Equal(t, want, got)
	})

	t.Run("Fit content table with title", func(t *testing.T) {

		tab, _ := NewTable(
			"TEST",
			[]string{"id", "key", "value"},
			[][]interface{}{{1, 2, 3}, {"test1", "test2", "test3"}, {3.14, 0.3, 1.0000003}},
		)
		tab.Options.SetCellFitContent(true)
		got := tab.Draw()

		want := `              TEST              
#======#=========#=============#
#  id  #   key   #    value    #
#======#=========#=============#
|  1   |  test1  |    3.14     |
+------+---------+-------------+
|  2   |  test2  |     0.3     |
+------+---------+-------------+
|  3   |  test3  |  1.0000003  |
+------+---------+-------------+
`
		assert.Equal(t, want, got)
	})
}

func TestMultiLineTable(t *testing.T) {
	t.Run("Multi line table", func(t *testing.T) {
		tab, _ := NewTable(
			"",
			[]string{"id", "value"},
			[][]interface{}{{1, 2, 3}, {"testtesttesttesttesttest", "testtesttest", "testtest"}},
		)
		tab.Options.SetCellLength(8)

		expectedTable := Table{
			Title:        "",
			headers:      []string{"id", "value"},
			columns:      [][]string{{"1", "2", "3"}, {"testtesttesttesttesttest", "testtesttest", "testtest"}},
			maxColLength: []int{2, 24},
			maxRowLength: []int{5, 24, 12, 8},
			Options: Options{
				titleAlign:     "center",
				headerAlign:    "center",
				cellLength:     8,
				cellPadding:    2,
				cellAlign:      "center",
				crossHeaderSym: '#',
				vHeaderSym:     '#',
				hHeaderSym:     '=',
				crossLineSym:   '+',
				vLineSym:       '|',
				hLineSym:       '-',
			},
		}
		assert.Equal(t, expectedTable, tab)
		got := tab.Draw()
		want := `#============#============#
#     id     #   value    #
#============#============#
|            |  testtest  |
|     1      |  testtest  |
|            |  testtest  |
+------------+------------+
|     2      |  testtest  |
|            |    test    |
+------------+------------+
|     3      |  testtest  |
+------------+------------+
`
		assert.Equal(t, want, got)
	})
	t.Run("Multi line table", func(t *testing.T) {
		tab, _ := NewTable(
			"",
			[]string{"id", "valuevalue"},
			[][]interface{}{{1, 2}, {"test", "test"}},
		)
		tab.Options.SetCellLength(8)
		got := tab.Draw()
		want := `#============#============#
#     id     #  valueval  #
#            #     ue     #
#============#============#
|     1      |    test    |
+------------+------------+
|     2      |    test    |
+------------+------------+
`
		assert.Equal(t, want, got)
	})

}
