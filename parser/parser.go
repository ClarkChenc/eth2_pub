package parser

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func IsDir(path string) bool {
	d, err := os.Stat(path)
	if err != nil {
		return false
	}

	return d.IsDir()
}

func TestParser() {
	cmds, err := os.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println("cmd: ", cmd)
		pkgdir := filepath.Join("./", cmd.Name())
		if !IsDir(pkgdir) {
			continue
		}
		pkgs, err := parser.ParseDir(token.NewFileSet(), pkgdir, nil, parser.ImportsOnly)
		if err != nil {
			log.Fatal(err)
		}

		for pkg := range pkgs {
			fmt.Println("package: ", pkg)
		}
	}
}
