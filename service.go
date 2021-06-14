package tuugen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type serviceEndpoint struct {
	name string
	args []string
}

func GenerateService(filepath string) error {

	//  load service template
	// t, err := template.ParseFiles("templates/service.go.tmpl")
	// if err != nil {
	// 	return err
	// }

	//	read interface from project built protos.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		return err
	}
	// generate the service template and store it to a file.
	err = ast.Print(fset, f)
	if err != nil {
		return err
	}
	return nil
}
