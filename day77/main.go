package main

import (
	"container/list"
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	lock sync.RWMutex
	Db   *sql.DB
}

type Table struct {
	TableName string `sql:"table_name"`
}

type CsvFields struct {
	Field   string `csv:"Field"`
	Type    string `csv:"type"`
	Null    string `csv:"es_null"`
	Key     string `csv:"key"`
	Default string `csv:"default"`
	Extra   string `csv:"extra"`
}

var db *Database
var wg sync.WaitGroup

func createMYSQLConnection(_ context.Context) (*Database, error) {
	dbPointer, err := sql.Open("mysql", "root:root@/data")
	if err != nil {
		return nil, err
	}

	db = &Database{lock: sync.RWMutex{}, Db: dbPointer}

	return db, nil
}

func (db *Database) GetTablesWithRows(ctx context.Context) (*list.List, error) {
	db.lock.RLock()
	ctx, cancel := context.WithCancel(ctx)
	defer db.lock.RUnlock()
	defer cancel()

	query, err := db.Db.QueryContext(ctx, "SELECT table_name FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_ROWS > 0 AND table_schema = ?", "data")
	if err != nil {
		log.Printf("Query table information error %v", err)
		return nil, err
	}
	defer query.Close()

	l := list.New()
	for query.Next() {
		var TableName string
		_ = query.Scan(&TableName)
		l.PushBack(TableName)
	}

	return l, nil
}

func ProcessTable(tableName chan string, ctx context.Context) {
	for {
		select {
		case table := <-tableName:
			db.lock.RLock()
			query, err := db.Db.QueryContext(ctx, fmt.Sprintf("SHOW COLUMNS FROM %s", table))
			queryRow, _ := db.Db.QueryContext(ctx, fmt.Sprintf("SELECT * FROM  %s LIMIT 1", table))

			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			db.lock.RUnlock()

			f, err := os.Create(fmt.Sprintf("backup/%s.csv", table))
			if err != nil {
				panic(err)
			}

			columns := []CsvFields{}
			for query.Next() {
				var Field, Type, Null, Extra, Key string
				var Default *string
				err := query.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
				if err != nil {
					panic(err)
				}
				columns = append(columns, CsvFields{Field: Field, Type: Type, Null: Null, Key: Key, Extra: Extra})
			}

			file := csv.NewWriter(f)
			_ = file.Write([]string{"Campo", "Tipo", "Null", "Primaria", "Default", "Extra"})
			for _, record := range columns {
				err = file.Write([]string{record.Field, record.Type, record.Null, record.Key, record.Default, record.Extra})

				if err != nil {
					panic(err)
				}
			}
			file.Flush()
			f.Close()
			query.Close()
			queryRow.Close()
		case <-ctx.Done():
			//
		}
	}

	wg.Done()
}

func RowsToMaps(rows *sql.Rows) ([]map[string]string, error) {
	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}
	columnCount := len(columns)

	cursor := make([]interface{}, columnCount)
	for i := 0; i < columnCount; i++ {
		var k string
		cursor[i] = &k
	}

	var resultMaps []map[string]string
	for rows.Next() {
		err := rows.Scan(cursor...)
		if err != nil {
			panic(err)
		}
		rowMap := make(map[string]string, columnCount)
		for i, columnPtr := range cursor {
			key := columns[i]
			var columnStr string
			if columnStrPtr := columnPtr.(*string); columnStrPtr != nil {
				columnStr = *columnStrPtr
			}
			rowMap[key] = columnStr
		}
		resultMaps = append(resultMaps, rowMap)
	}
	if err := rows.Err(); err != nil {
		return resultMaps, err
	}
	return resultMaps, nil
}

func OutputBackup(ctx context.Context, l *list.List) {
	ctx, cancel := context.WithCancel(ctx)
	table := make(chan string)

	if err := os.Mkdir("backup", os.ModePerm); err != nil {
		if errors.Is(err, os.ErrExist) {
			os.Remove("backup")
		} else {
			panic(err)
		}

		_ = os.Mkdir("backup", os.ModePerm)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ProcessTable(table, ctx)
	}

	tableName := l.Front()
	for tableName != nil {
		s := tableName.Value.(string)
		table <- s
		tableName = tableName.Next()
	}

	fmt.Printf("Gorutinas running: %d\n", runtime.NumGoroutine())
	fmt.Printf("CPUs using: %v\n", runtime.NumCPU())
	fmt.Printf("Num nodes: %d\n", l.Len())
	cancel()
	wg.Wait()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	ctx := context.Background()
	db, err := createMYSQLConnection(ctx)
	if err != nil {
		log.Fatalf("Opendatabase error %v", err)
		return
	}
	defer db.Db.Close()

	l, err := db.GetTablesWithRows(ctx)
	if err != nil {
		log.Fatalf("Error query: %v", err)
	}
	OutputBackup(ctx, l)

}
