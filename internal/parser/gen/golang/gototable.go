//Copyright 2013 Vastech SA (PTY) LTD
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package golang

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"path"
	"text/template"

	"github.com/aggronmagi/gocc/internal/io"
	"github.com/aggronmagi/gocc/internal/parser/lr1/items"
	"github.com/aggronmagi/gocc/internal/parser/symbols"
)

type gotoTableData struct {
	NumNTSymbols int
	Rows         [][]gotoRowElement
}

type gotoRowElement struct {
	NT    string
	State int
	Pad   int
}

func GenGotoTable(outDir string, itemSets *items.ItemSets, sym *symbols.Symbols, zip bool) {
	if zip {
		GenCompGotoTable(outDir, itemSets, sym)
		return
	}
	tmpl, err := template.New("parser goto table").Parse(gotoTableSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, getGotoTableData(itemSets, sym)); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(outDir, "parser", "gototable.go"), wr.Bytes())
}

func getGotoTableData(itemSets *items.ItemSets, sym *symbols.Symbols) *gotoTableData {
	data := &gotoTableData{
		NumNTSymbols: sym.NumNTSymbols(),
		Rows:         make([][]gotoRowElement, itemSets.Size()),
	}
	for i, set := range itemSets.List() {
		data.Rows[i] = getGotoRowData(set, sym)
	}
	return data
}

func getGotoRowData(itemSet *items.ItemSet, sym *symbols.Symbols) []gotoRowElement {
	row := make([]gotoRowElement, sym.NumNTSymbols())
	var max int
	for i, nt := range sym.NTList() {
		row[i].NT = nt
		row[i].State = itemSet.NextSetIndex(nt)
		n := nbytes(row[i].State)
		if n > max {
			max = n
		}
	}
	// Calculate padding.
	for i := range row {
		row[i].Pad = max + 1 - nbytes(row[i].State)
	}
	return row
}

const gotoTableSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser

const numNTSymbols = {{.NumNTSymbols}}

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{
{{- range $i, $r := .Rows }}
	gotoRow{ // S{{$i}}
	{{- range $j, $gto := . }}
		{{ printf "%d,%*c// %s" $gto.State $gto.Pad ' ' $gto.NT }}
	{{- end }}
	},
{{- end }}
}
`

func genEnc(v interface{}) string {
	buf := bytes.NewBuffer(nil)
	z := gzip.NewWriter(buf)
	enc := gob.NewEncoder(z)
	if err := enc.Encode(v); err != nil {
		panic(err)
	}
	if err := z.Close(); err != nil {
		panic(err)
	}
	strBuf := bytes.NewBuffer(nil)
	fmt.Fprintf(strBuf, "[]byte{\n")
	b := buf.Bytes()
	for len(b) > 0 {
		n := 16
		if n > len(b) {
			n = len(b)
		}
		fmt.Fprintf(strBuf, "\t\t")
		for i, c := range b[:n] {
			if i != 0 {
				strBuf.WriteString(" ")
			}
			fmt.Fprintf(strBuf, "0x%02x,", c)
		}
		fmt.Fprintf(strBuf, "\n")
		b = b[n:]
	}
	fmt.Fprintf(strBuf, "\t}")
	return string(strBuf.Bytes())
}

func GenCompGotoTable(outDir string, itemSets *items.ItemSets, sym *symbols.Symbols) {
	numNTSymbols := sym.NumNTSymbols()
	rows := make([][]int, itemSets.Size())
	for i, set := range itemSets.List() {
		rows[i] = make([]int, numNTSymbols)
		for j, nt := range sym.NTList() {
			rows[i][j] = set.NextSetIndex(nt)
		}
	}
	bytesStr := genEnc(rows)
	v := struct {
		NumNTSymbols int
		Bytes        string
	}{
		NumNTSymbols: numNTSymbols,
		Bytes:        bytesStr,
	}
	tmpl, err := template.New("parser goto table").Parse(gotoTableCompSrc[1:])
	if err != nil {
		panic(err)
	}
	wr := new(bytes.Buffer)
	if err := tmpl.Execute(wr, v); err != nil {
		panic(err)
	}
	io.WriteFile(path.Join(outDir, "parser", "gototable.go"), wr.Bytes())
}

const gotoTableCompSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
)

const numNTSymbols = {{.NumNTSymbols}}

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{}

func init() {
	tab := [][]int{}
	data := {{.Bytes}}
	buf, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&tab); err != nil {
		panic(err)
	}
	for i := 0; i < numStates; i++ {
		for j := 0; j < numNTSymbols; j++ {
			gotoTab[i][j] = tab[i][j]
		}
	}
}
`
