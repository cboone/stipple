package draw

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/cboone/brodot/canvas"
)

var visualFlag = flag.Bool("visual", false, "print visual output")
var updateFlag = flag.Bool("update", false, "update golden files")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func printVisual(t *testing.T, name string, c *canvas.Canvas) {
	if *visualFlag {
		t.Logf("\n=== %s ===\n%s", name, c.Frame())
	}
}

func goldenPath(name string) string {
	return filepath.Join("testdata", name+".golden")
}

func assertGolden(t *testing.T, name string, c *canvas.Canvas) {
	t.Helper()

	actual := c.Frame()
	path := goldenPath(name)

	if *updateFlag {
		if err := os.MkdirAll("testdata", 0755); err != nil {
			t.Fatalf("failed to create testdata directory: %v", err)
		}
		if err := os.WriteFile(path, []byte(actual), 0644); err != nil {
			t.Fatalf("failed to write golden file %s: %v", path, err)
		}
		t.Logf("updated golden file: %s", path)
		return
	}

	expected, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read golden file %s (run with -update to create): %v", path, err)
	}

	if actual != string(expected) {
		t.Errorf("output does not match golden file %s\n--- expected ---\n%s\n--- actual ---\n%s",
			path, string(expected), actual)
	}
}
