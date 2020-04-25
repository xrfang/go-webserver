package main

import (
	"fmt"
	"os"
	"path/filepath"

	res "github.com/xrfang/go-res"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	PROJ_ROOT, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "/..")
	assert(os.Chdir(PROJ_ROOT))
	fns, err := filepath.Glob("bin/*")
	assert(err)
	root := filepath.Join(PROJ_ROOT, "resources")
	for _, fn := range fns {
		fmt.Printf("pack: processing %s...\n", fn)
		assert(res.Pack(root, fn))
	}
	fmt.Printf("pack: processed %d files\n", len(fns))
}
