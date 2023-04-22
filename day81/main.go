package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"strings"
)

type DirInput string
type DirOutput string
type DirName string

type Ignore struct {
	dirs []DirName
}

type Move struct {
	DirIn                 DirInput
	DirOut                DirOutput
	Containers            map[string][]DirName
	fsys                  fs.FS
	IgnoringNonContainers bool
	*Ignore
}

type Path interface {
	ConcatFile(string, *string) *string
}

func NewMove(in DirInput, out DirOutput, containers map[string][]DirName, inc bool, ig *Ignore) *Move {
	return &Move{
		DirIn:                 in,
		DirOut:                out,
		Containers:            containers,
		fsys:                  os.DirFS(string(in)),
		IgnoringNonContainers: inc,
		Ignore:                ig,
	}
}

func NewIgnore(dirs string) *Ignore {
	ignore := strings.Split(dirs, ",")
	var dirNames []DirName
	for _, v := range ignore {
		dirNames = append(dirNames, DirName(strings.ToLower(v)))
	}

	return &Ignore{dirs: dirNames}
}

func (i *Ignore) Empty() bool {
	return len(i.dirs) == 0
}

func (i *Ignore) IsIgnored(dirname DirName) bool {

	if i.Empty() {
		return false
	}

	dn := strings.ToLower(string(dirname))
	for _, dir := range i.dirs {
		if strings.Contains(string(dir), dn) {
			return true
		}
	}

	return false
}

func TrimPath(path string) string {
	pout := path
	if strings.HasSuffix(path, "/") {
		pout = strings.TrimSuffix(path, "/")
	}

	return pout
}

func (dn *DirInput) ConcatFile(fn string, _ *string) *string {
	d := string(*dn)
	var dout string = TrimPath(d)
	dout = fmt.Sprintf("%s/%s", dout, fn)

	return &dout
}

func (dn *DirOutput) ConcatFile(fn string, container *string) *string {
	d := string(*dn)
	containerDir := ""
	dout := TrimPath(d)

	if container != nil && (*container) != "" {
		containerDir = *container
		dout = fmt.Sprintf("%s/%s/%s", dout, containerDir, fn)
	} else {
		dout = fmt.Sprintf("%s/%s", dout, fn)
	}

	return &dout
}

func CreateIfNotExists(dn Path, container *string) (bool, error) {
	fullPath := dn.ConcatFile(*container, nil)
	_ = os.Mkdir(*fullPath, os.ModePerm)
	return true, nil
}

func getContainer(cons map[string][]DirName, extension string) string {
	for k, v := range cons {
		for _, vv := range v {
			if strings.ToLower(string(vv)) == extension {
				return string(k)
			}
		}
	}

	return ""
}

func (m *Move) Walk(removeAfterMove bool) {

	fs.WalkDir(m.fsys, ".", func(path string, d fs.DirEntry, err error) error {
		var fileExtension string = ""
		var to *string

		if d.IsDir() {
			return nil
		}
		finfo, errInfo := d.Info()
		if errInfo != nil {
			log.Printf("File error: %v", errInfo)
			return errInfo
		}

		extension := strings.Split(finfo.Name(), ".")

		if m.IsIgnored(DirName(finfo.Name())) {
			log.Printf("Archivo ignorado: %s\n", extension[0])
			return nil
		}

		if len(extension) > 1 {
			fileExtension = extension[len(extension)-1]
		} else {
			fileExtension = finfo.Name()
		}

		container := getContainer(m.Containers, strings.ToLower(fileExtension))
		from := m.DirIn.ConcatFile(path, nil)

		if len(container) > 0 {
			to = m.DirOut.ConcatFile(finfo.Name(), &container)
		} else {
			if m.IgnoringNonContainers {
				return nil
			}
			to = m.DirOut.ConcatFile(finfo.Name(), nil)
		}

		CreateIfNotExists(&m.DirOut, &container)
		to2 := strings.ReplaceAll(*to, " ", "_")
		if removeAfterMove {
			errMove := exec.Command("mv", *from, to2).Run()
			if errMove != nil {
				log.Fatal("Err", errMove)
				return errMove
			}
			fmt.Printf("Deleting: %s\n", *from)
		} else {
			cmd := exec.Command("cp", (*from), to2)
			err = cmd.Run()
			if err != nil {
				log.Fatal("cp ", *from, " ", to2, "\n", err.Error())
			}
		}
		return nil
	})
}

func main() {
	var dirIn, dirOut string
	ignoreNoConteinerizeds := true
	removeAfterMove := false
	var ignoreDirs string
	dirsContents := make(map[string][]DirName)

	flag.StringVar(&dirIn, "dir-in", "", "Define directorio de entrada")
	flag.StringVar(&dirOut, "dir-out", "", "Define directorio de salida")
	flag.BoolVar(&ignoreNoConteinerizeds, "ignore-no-containerizeds", true, "-ignore-no-containerizeds=1")
	flag.BoolVar(&removeAfterMove, "autoremove", false, "Indica que se desea eliminar del origen despues de ser movido el archivo")
	flag.StringVar(&ignoreDirs, "ignore-dirs", "", "Indica una lista de directorios separada por coma (,) los nombres de los directorios en este parametro no seran tocados")

	flag.Func("segregate-dirs", "Indica la lista de directorios de salida, usa el formato dir-name:extension", func(in string) error {
		if len(strings.TrimRight(in, " ")) == 0 {
			return errors.New("segregate-dirs should be not empty")
		}

		listDir := strings.Split(in, ",")
		for _, d := range listDir {
			parts := strings.Split(d, ":")
			dirsContents[parts[0]] = append(dirsContents[parts[0]], DirName(strings.ToLower(parts[1])))
		}
		return nil
	})

	flag.Parse()

	move := NewMove(DirInput(dirIn), DirOutput(dirOut), dirsContents, ignoreNoConteinerizeds, NewIgnore(ignoreDirs))
	move.Walk(removeAfterMove)
}
