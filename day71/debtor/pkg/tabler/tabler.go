package tabler

import "fmt"

type CornerCharacter string
type LineCharacter string
type ColumnNames []string

var tableSpaceFactor int = 4

const MaxCharactesSpaces = 20
const MinCharactersSpace = 20

type Table struct {
	Corner      CornerCharacter
	Line        LineCharacter
	ColumnNames ColumnNames
	Row         []Row
}

type Row struct {
	Coumns []Column
	Id     int
}

type Column struct {
	Id    int
	Value string
}

// New creates a new table object
func New(corner CornerCharacter, line LineCharacter) *Table {
	return &Table{
		Corner: corner,
		Line:   line,
	}
}

func (t *Table) AddRow(row *Row) *Table {
	t.Row = append(t.Row, *row)

	for _, col := range row.Coumns {
		if aux := len(col.Value); aux > tableSpaceFactor {
			tableSpaceFactor = aux
		}
	}

	return t
}

// SetCorner sets a corner character
func (t *Table) SetCorner(corner CornerCharacter) *Table {
	t.Corner = corner
	return t
}

// SetLineCharacter sets the underline and overline between rows
func (t *Table) SetLineCharacter(line LineCharacter) *Table {
	t.Line = line
	return t
}

func (t *Table) SetColumnsName(columns ColumnNames) *Table {
	t.ColumnNames = columns

	for _, col := range columns {
		if aux := len(col); aux > tableSpaceFactor {
			tableSpaceFactor = aux
		}
	}

	if tableSpaceFactor > MaxCharactesSpaces {
		tableSpaceFactor = 20
	} else if tableSpaceFactor < MinCharactersSpace {
		tableSpaceFactor = MinCharactersSpace
	}

	return t
}

// columnsNameLarge gets the much large column name string as int
// and return it
func (t *Table) ColumnNameLarge() int {
	muchLargeColumnName := 0

	for _, columnName := range t.ColumnNames {
		if aux := len(columnName); aux > muchLargeColumnName {
			muchLargeColumnName = aux
		}
	}

	if tableSpaceFactor > MaxCharactesSpaces {
		tableSpaceFactor = 20
	} else if tableSpaceFactor < MinCharactersSpace {
		tableSpaceFactor = MinCharactersSpace
	}

	return muchLargeColumnName
}

func (t *Table) printInterlineRow() {
	for _, colNameInfo := range t.ColumnNames {
		fmt.Print(t.Corner)
		for i := 0; i < len(colNameInfo)+tableSpaceFactor; i++ {
			fmt.Print(t.Line)
		}
	}

	fmt.Println(t.Corner)
}

// PrintColumns draws columns headers
func (t *Table) printColumns() *Table {

	t.printInterlineRow()

	for _, colName := range t.ColumnNames {
		fmt.Print("|")
		fmt.Printf(" %s", colName)
		for i := 0; i < tableSpaceFactor-1; i++ {
			fmt.Print(" ")
		}

	}

	fmt.Println("|")
	t.printInterlineRow()
	return t
}

// PrintRows prints row-by-row all rows in the table
func (t *Table) PrintRows() {
	rows := t.Row
	var colNameLenghts []int

	for _, colName := range t.ColumnNames {
		colNameLenghts = append(colNameLenghts, len(colName)+tableSpaceFactor)
	}

	for _, row := range rows {
		columns := row.Coumns
		for colIndex, col := range columns {
			spaces := ""
			cellValue := col.Value

			if len(cellValue) >= tableSpaceFactor {
				cellValue = fmt.Sprintf("%s...", cellValue[:MaxCharactesSpaces-3])
			}

			// It calcs amount spaces between cells and column divider
			lenSpaces := (colNameLenghts[colIndex] - len(cellValue)) - 1
			for i := 0; i < lenSpaces; i++ {
				spaces = fmt.Sprintf("%s ", spaces)
			}
			fmt.Print("| ", cellValue, spaces)

		}
		fmt.Println("|")
	}
}

// Prints table
func (t *Table) Show() {
	t.printColumns()
	t.PrintRows()
	t.printInterlineRow()
}
