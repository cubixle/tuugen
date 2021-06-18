package main

import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"path/filepath"
	"strings"
)

func GenerateService(FS embed.FS, cfg Config) error {
	return createFileFromProto(FS, cfg, "templates/service.go.tmpl", "./internal/service/service.go")
}

func GenerateInteractor(FS embed.FS, cfg Config) error {
	return createFileFromProto(FS, cfg, "templates/interactors.go.tmpl", "./internal/interactors/interactors.go")
}

type serviceDef struct {
	Imports []string
	Funcs   []serviceFuncDef
}

type serviceFuncDef struct {
	Name    string
	Args    []string
	Returns []string
}

func (s serviceFuncDef) ListToStr(list []string) string {
	return strings.Join(list, ", ")
}

func (s serviceFuncDef) ArgsToStr() string {
	args := []string{}
	for _, a := range s.Args {
		switch a {
		case "context.Context":
			args = append(args, fmt.Sprintf("ctx %s", a))
		default:
			args = append(args, fmt.Sprintf("req %s", a))
		}
	}
	return s.ListToStr(args)
}

// generate the service/interactor template and store it to a file.
func createFileFromProto(FS embed.FS, cfg Config, templateFile, outputFile string) error {
	def, err := parseProto(cfg)
	if err != nil {
		return err
	}

	// dirty check to see if we need to import the interactors.
	if strings.Contains(templateFile, "service") {
		def.Imports = append(def.Imports, cfg.ImportPath+"/internal/interactors")
	}

	f, err := createFile(outputFile)
	if err != nil {
		return err
	}

	t, err := template.ParseFS(FS, templateFile)
	if err != nil {
		return err
	}

	if err := t.Execute(f, def); err != nil {
		return err
	}

	return f.Close()
}

func parseProto(cfg Config) (serviceDef, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, cfg.GRPCFile, nil, 0)
	if err != nil {
		return serviceDef{}, fmt.Errorf("failed to parse compiled proto go file: %w", err)
	}

	imports := parseImports(cfg, f)
	funcs := parseFuncs(cfg, f)

	return serviceDef{Funcs: funcs, Imports: imports}, nil
}

func parseImports(cfg Config, f *ast.File) []string {
	imports := []string{}
	servicePath := filepath.Dir(cfg.GRPCFile)
	for _, i := range f.Imports {
		imports = append(imports, strings.Replace(i.Path.Value, `"`, "", -1))
	}
	imports = append(imports, strings.Join([]string{cfg.ImportPath, strings.TrimLeft(servicePath, "/")}, "/"))
	return imports
}

func parseFuncs(cfg Config, f *ast.File) []serviceFuncDef {
	funcs := []serviceFuncDef{}
	for _, m := range f.Scope.Objects {
		if m.Name != cfg.ServiceName+"Server" {
			continue
		}
		decl := m.Decl.(*ast.TypeSpec)
		t := decl.Type.(*ast.InterfaceType)

		for _, m := range t.Methods.List {
			name := m.Names[0].Name
			if strings.Contains(name, "mustEmbed") {
				continue
			}
			f := serviceFuncDef{Name: name}
			ft := m.Type.(*ast.FuncType)
			if ft.Params != nil {
				f.Args = parseFieldList(ft.Params)
			}

			if ft.Results != nil {
				f.Returns = parseFieldList(ft.Results)
			}
			funcs = append(funcs, f)
		}
	}
	return funcs
}

func parseFieldList(fl *ast.FieldList) []string {
	args := []string{}
	if fl.List == nil {
		return args
	}
	for _, a := range fl.List {
		argStr := ""
		switch a.Type.(type) {
		case *ast.SelectorExpr:
			as := a.Type.(*ast.SelectorExpr)
			argStr = as.X.(*ast.Ident).Name + "." + as.Sel.Name
		case *ast.StarExpr:
			as := a.Type.(*ast.StarExpr)
			argStr = "*service." + as.X.(*ast.Ident).Name
		case *ast.Ident:
			argStr = a.Type.(*ast.Ident).Name
		}
		args = append(args, argStr)
	}
	return args
}
