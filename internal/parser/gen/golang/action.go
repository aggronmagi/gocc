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
	"path"

	"github.com/aggronmagi/gocc/internal/io"
)

func GenAction(outDir string) {
	io.WriteFileString(path.Join(outDir, "parser", "action.go"), actionSrc[1:])
}

const actionSrc = `
// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"fmt"
)

type action interface {
	act()
	String() string
}

type (
	accept bool
	shift  int // value is next state index
	reduce int // value is production index
)

func (this accept) act() {}
func (this shift) act()  {}
func (this reduce) act() {}

func (this accept) Equal(that action) bool {
	if _, ok := that.(accept); ok {
		return true
	}
	return false
}

func (this reduce) Equal(that action) bool {
	that1, ok := that.(reduce)
	if !ok {
		return false
	}
	return this == that1
}

func (this shift) Equal(that action) bool {
	that1, ok := that.(shift)
	if !ok {
		return false
	}
	return this == that1
}

func (this accept) String() string { return "accept(0)" }
func (this shift) String() string  { return fmt.Sprintf("shift:%d", this) }
func (this reduce) String() string {
	return fmt.Sprintf("reduce:%d(%s)", this, productionsTable[this].String)
}
`
