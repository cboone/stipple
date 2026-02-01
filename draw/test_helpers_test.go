package draw

import (
	"flag"
	"os"
	"testing"

	"github.com/cboone/brodot/canvas"
)

var visualFlag = flag.Bool("visual", false, "print visual output")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func printVisual(t *testing.T, name string, c *canvas.Canvas) {
	if *visualFlag {
		t.Logf("\n=== %s ===\n%s", name, c.Frame())
	}
}
