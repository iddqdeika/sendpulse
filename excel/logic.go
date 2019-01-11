package excel

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gopkg.in/cheggaaa/pb.v1"
)

//structured object for simplier interpreting data for writing excel sheet.
//has list of columns with their nums
type Table struct {
	Columns map[string]int
	Rows    []Row
}

//substruct of table, contains cells. key of map is num if column
type Row struct {
	Cells map[int]string
}

//set cell value for current row.
func (t *Table) SetCellValue(column string, value string) {
	if t.Rows==nil{
		t.AddRow()
	}
	if !t.ContainsColumn(column) {
		t.AddColumn(column)
	}
	t.Rows[len(t.Rows)-1].Cells[t.Columns[column]] = value
}

//check has table given column yet or not
func (t *Table) ContainsColumn(column string) bool {
	_, ok := t.Columns[column]
	if ok {
		return true
	}
	return false
}

//add column to table
func (t *Table) AddColumn(column string) {
	if t.Columns == nil {
		t.Columns = make(map[string]int)
	}
	t.Columns[column] = len(t.Columns)
}

//add new row to table
func (t *Table) AddRow() error {
	if len(t.Rows) < 1048576 {
		row := Row{}
		row.Cells = make(map[int]string)
		t.Rows = append(t.Rows, row)
		return nil
	}
	return errors.New("Not wnough rows count.")
}

//write Table object content to given xlsx file into sheet with given name
func WriteTable(xlsx *excelize.File, sheetname string, table *Table) {
	fmt.Println("\t" + "writing sheet \"" + sheetname + "\"...")
	xlsx.NewSheet(sheetname)
	for k, v := range table.Columns {
		columnname := getColumnName(v)
		xlsx.SetCellValue(sheetname, columnname+"1", k)
	}
	var i int

	bar := pb.StartNew(len(table.Rows))
	for k, v := range table.Rows {
		i++
		bar.Increment()
		rowname := strconv.Itoa(k + 2)

		//xlsx.SetSheetRow(sheetname, "A" + rowname,&sl)
		for kk, vv := range v.Cells {
			columnname := getColumnName(kk)
			xlsx.SetCellValue(sheetname, columnname+rowname, vv)
		}

	}
	bar.Finish()
}

//function to get column name by column number. supports up to 17526 columns
func getColumnName(v int) string {
	var columnname string
	if v <= 17526 {
		alfabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		columnnum := v
		if columnnum >= len(alfabet) {
			if columnnum >= len(alfabet)*len(alfabet) {
				first := (columnnum - (columnnum % (len(alfabet) * len(alfabet)))) / (len(alfabet) * len(alfabet))
				last := (columnnum - first) % len(alfabet)
				mid := (columnnum - (first * len(alfabet) * len(alfabet)) - last) / len(alfabet)
				columnname = alfabet[first] + alfabet[mid] + alfabet[last]
			} else {
				last := alfabet[columnnum%len(alfabet)]
				first := alfabet[columnnum/len(alfabet)-1]
				columnname = first + last
			}
		} else {
			first := alfabet[columnnum]
			columnname = first
		}
	} else {
		columnname = getColumnName(17526)
	}
	return columnname
}
