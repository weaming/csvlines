package csvlines

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

type CSVLines struct {
	sync.Mutex
	Path   string
	File   *os.File
	Writer *csv.Writer
}

func New(filepath string) *CSVLines {
	return &CSVLines{Path: filepath}
}

func (r *CSVLines) Init(truncate bool) {
	mode := os.O_WRONLY | os.O_CREATE
	if truncate {
		mode |= os.O_TRUNC
	} else {
		mode |= os.O_APPEND
	}

	file, e := os.OpenFile(r.Path, mode, 0644)
	CheckError("OpenFile", e)

	r.File = file
	r.Writer = csv.NewWriter(file)
}

func (r *CSVLines) Write(vs []string) {
	r.Lock()
	defer r.Unlock()
	defer r.Writer.Flush()

	e := r.Writer.Write(vs)
	CheckError("Write", e)
}

func (r *CSVLines) WriteAll(vss [][]string) {
	r.Lock()
	defer r.Unlock()
	defer r.Writer.Flush()

	e := r.Writer.WriteAll(vss)
	CheckError("Writeall", e)
}

func (r *CSVLines) Close() {
	e := r.File.Close()
	CheckError("Close", e)
}

func Str(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func CheckError(message string, e error) {
	if e != nil {
		log.Fatal(fmt.Sprintf("%s: %v", message, e))
	}
}
