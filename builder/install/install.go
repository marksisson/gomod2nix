package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"os/exec"
	"strconv"
)

const filename = "tools.go"

func main() {
	fset := token.NewFileSet()

	var src []byte
	{
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		src, err = io.ReadAll(f)
		if err != nil {
			panic(err)
		}
	}

	f, err := parser.ParseFile(fset, filename, src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range f.Imports {
		path, err := strconv.Unquote(s.Path.Value)
		if err != nil {
			panic(err)
		}

    pathWithVersion := path + "@latest"

    varName := "GOPATH"
    value := os.Getenv(varName)
    if value != "" {
      fmt.Printf("%s=%s\n", varName, value)
    } else {
      fmt.Printf("%s is not set\n", varName)
    }
    varName2 := "PWD"
    value2 := os.Getenv(varName2)
    if value != "" {
      fmt.Printf("%s=%s\n", varName2, value2)
    } else {
      fmt.Printf("%s is not set\n", varName2)
    }

		cmd := exec.Command("go", "install", "-v", pathWithVersion)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
    cmd.Env = os.Environ()

		fmt.Printf("Executing '%s'\n", cmd)

		err = cmd.Start()
		if err != nil {
			panic(err)
		}

		err = cmd.Wait()
		if err != nil {
			panic(err)
		}
	}
}
