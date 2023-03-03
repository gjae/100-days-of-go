package main

import (
	"github.com/gjae/tabler/tablerer"
)

func main() {
	tab := tablerer.New(tablerer.CornerCharacter("+"), tablerer.LineCharacter("-"))
	columns := tablerer.ColumnNames([]string{"First", "Second", "Third", "Fourth", "Fifth"})
	newColumns := []tablerer.Column{
		tablerer.Column{Id: 1, Value: "Primero"},
		tablerer.Column{Id: 2, Value: "Segundo"},
		tablerer.Column{Id: 3, Value: "Tercero"},
		tablerer.Column{Id: 4, Value: "Cuarto"},
		tablerer.Column{Id: 5, Value: "Quinto"},
	}
	newColumns2 := []tablerer.Column{
		tablerer.Column{Id: 6, Value: "1 Primero"},
		tablerer.Column{Id: 7, Value: "2 Segundo"},
		tablerer.Column{Id: 8, Value: "3 Tercero"},
		tablerer.Column{Id: 9, Value: "4 Cuarto"},
		tablerer.Column{Id: 10, Value: "Quinto"},
	}

	tab.SetColumnsName(columns)

	tab.AddRow(&tablerer.Row{Id: 1, Coumns: newColumns}).AddRow(&tablerer.Row{Id: 2, Coumns: newColumns2})

	tab.Show()
}
