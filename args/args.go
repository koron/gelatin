package args

import (
	"os"

	"github.com/koron/gelatin/omap"
)

// Root is root mode.
var Root = &Mode{
	subModes: omap.New(),
}

// Parse aguments.
func Parse() error {
	return Root.Parse(os.Args...)
}
