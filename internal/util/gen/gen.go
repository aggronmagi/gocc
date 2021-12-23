package gen

import (
	"path"

	"github.com/aggronmagi/gocc/internal/util/gen/golang"
)

func Gen(outDir string) {
	outDir = path.Dir(outDir)
	golang.GenRune(outDir)
	golang.GenLitConv(outDir)
}
