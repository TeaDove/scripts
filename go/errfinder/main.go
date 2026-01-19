package main

import (
	"fmt"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

const (
	literalErr = "err"
	literalNil = "nil"
)

func isSimple(ifStmt *ast.IfStmt) bool {
	if len(ifStmt.Body.List) != 1 {
		return false
	}

	if len(ifStmt.Body.List) != 1 {
		return false
	}

	ret, ok := ifStmt.Body.List[0].(*ast.ReturnStmt)
	if !ok {
		return false
	}

	if len(ret.Results) == 0 {
		return false
	}

	last := ret.Results[len(ret.Results)-1]
	ident, ok := last.(*ast.Ident)
	if ok {
		if ident.Name == literalErr {
			return true
		}
	}

	call, ok := last.(*ast.CallExpr)
	if ok {
		for _, arg := range call.Args {
			ident, ok = arg.(*ast.Ident)
			if ok && ident.Name == literalErr {
				return true
			}
		}
	}

	return false
}

func analyzeFile(fset *token.FileSet, filename string) (int, int, error) {
	src, err := os.ReadFile(filename)
	if err != nil {
		return 0, 0, errors.Wrap(err, "reading source file")
	}

	f, err := parser.ParseFile(fset, filename, src, 0)
	if err != nil {
		return 0, 0, errors.Wrap(err, "parse file")
	}

	var total, simple int

	ast.Inspect(f, func(n ast.Node) bool {
		ifStmt, ok := n.(*ast.IfStmt)
		if !ok {
			return true
		}

		bin, ok := ifStmt.Cond.(*ast.BinaryExpr)
		if !ok || bin.Op != token.NEQ {
			return true
		}

		left, lOk := bin.X.(*ast.Ident)
		right, rOk := bin.Y.(*ast.Ident)
		if !(lOk && rOk && left.Name == literalErr && right.Name == literalNil) {
			return true
		}

		total++
		if isSimple(ifStmt) {
			fmt.Printf("Simple err:\n")
			simple++
		} else {
			fmt.Printf("Complex err:\n")
		}

		pos := fset.Position(ifStmt.Pos())
		fmt.Printf("%s:%d\n", pos.Filename, pos.Line)
		fmt.Printf("%s\n\n", getIfSnippet(fset, src, ifStmt))

		return true
	})

	return total, simple, nil
}

func getIfSnippet(fset *token.FileSet, src []byte, ifStmt *ast.IfStmt) string {
	startPos := fset.Position(ifStmt.Pos())
	endPos := fset.Position(ifStmt.End())

	// используем byte-Offsets для вытаскивания исходника
	start := startPos.Offset
	end := endPos.Offset

	if start < 0 {
		start = 0
	}
	if end > len(src) {
		end = len(src)
	}
	if start >= end {
		return ""
	}

	snippet := src[start:end]
	return string(snippet)
}

func analyzeFiles(path string) (int, int, error) {
	fset := token.NewFileSet()

	info, err := os.Stat(path)
	if err != nil {
		return 0, 0, errors.Wrap(err, "stat path")
	}

	if !info.IsDir() {
		return analyzeFile(fset, path)
	}

	var total, simple int
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(p) == ".go" {
			totalL, simpleL, err := analyzeFile(fset, p)
			if err != nil {
				return errors.Wrap(err, "analyze file")
			}

			total += totalL
			simple += simpleL
		}
		return nil
	})
	if err != nil {
		return 0, 0, errors.Wrap(err, "walk path")
	}

	return total, simple, nil
}

func main() {
	total, simple, err := analyzeFiles(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Всего обработок err: %d\n", total)
	fmt.Printf("Простых: %d\n", simple)
	fmt.Printf("Сложных: %d\n", total-simple)
}
