package goconsole

import (
	"fmt"
	"github.com/995933447/std-go/scan"
	"testing"
)

func TestRun(t *testing.T)  {
	Register("foo", "foo -bar $bar", func() {
		fmt.Println(scan.OptStr("bar"))
	})
	Run()
}
