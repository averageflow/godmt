package syntaxtree

import "go/ast"

const (
	MapType   = 1
	VarType   = 2
	ConstType = 3
)

type RawScannedType struct {
	Name         string
	Kind         string
	Value        interface{}
	Doc          []string
	InternalType int
}

type visitor int

var ScanResult []RawScannedType

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {

	case *ast.Ident:
		if d.Obj == nil {
			return v + 1
		}

		if d.Obj.Kind == ast.Typ {

		}

		if d.Obj.Kind == ast.Con || d.Obj.Kind == ast.Var {
			result := parseConstantsAndVariables(d)

			for i := range result {
				ScanResult = append(ScanResult, result[i])
			}
		}

		break
	}
	//fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}
