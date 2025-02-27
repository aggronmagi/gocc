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

package items

import (
	"fmt"

	"github.com/aggronmagi/gocc/internal/util"
)

type CharRange struct {
	From rune
	To   rune
}

func (this CharRange) String() string {
	return fmt.Sprintf("[%s,%s]", util.RuneToString(this.From), util.RuneToString(this.To))
}

func (this CharRange) Equal(that CharRange) bool {
	return this.From == that.From && this.To == that.To
}
