package table

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testing"
)

var ERR error
var countChars int

func benchmarkCreate(b *testing.B, buffer, filezise int) {
	filename := path.Join(os.TempDir(), strconv.Itoa(buffer))

	filename = filename + "-" + strconv.Itoa(filezise)

	var err error

	for i:= 0; i < b.N; i ++ {
		err = Create(filename, buffer,filezise)
	}

	ERR = err
}