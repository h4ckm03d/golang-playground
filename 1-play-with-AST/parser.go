package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

func resolveAssociate(model *Model, modelMap map[string]*Model, parents map[string]bool) {
	parents[model.Name] = true

	for i, field := range model.Fields {
		if field.Association != nil && field.Association.Type != AssociationNone {
			continue
		}

		str := strings.Trim(field.Type, "[]*")
		if modelMap[str] != nil && !parents[str] {
			resolveAssociate(modelMap[str], modelMap, parents)

			var assoc int
			switch string([]rune(field.Type)[0]) {
			case "[":
				if validateForeignKey(modelMap[str].Fields, model.Name) {
					assoc = AssociationHasMany
					break
				}
				assoc = AssociationBelongsTo
			default:
				if validateForeignKey(modelMap[str].Fields, model.Name) {
					assoc = AssociationHasOne
					break
				}
				assoc = AssociationBelongsTo
			}
			model.Fields[i].Association = &Association{Type: assoc, Model: modelMap[str]}
		} else {
			model.Fields[i].Association = &Association{Type: AssociationNone}
		}
	}
}

func validateForeignKey(fields []*Field, name string) bool {
	for _, field := range fields {
		val := name + "ID"
		if field.Name == val {
			return true
		}
	}
	return false
}

func parseField(field *ast.Field) (*Field, error) {
	if len(field.Names) != 1 {
		return nil, errors.New("Failed to read model files. Please fix struct.")
	}

	fieldName := field.Names[0].Name

	var fieldType string

	switch x := field.Type.(type) {
	case *ast.Ident: // e.g. string
		fieldType = x.Name

	case *ast.SelectorExpr: // e.g. time.Time, sql.NullString
		switch x2 := x.X.(type) {
		case *ast.Ident:
			fieldType = x2.Name + "." + x.Sel.Name
		}

	case *ast.ArrayType: // e.g. []Email
		switch x2 := x.Elt.(type) {
		case *ast.Ident:
			fieldType = "[]" + x2.Name

		case *ast.StarExpr: // e.g. []*Email
			switch x3 := x2.X.(type) {
			case *ast.Ident:
				fieldType = "[]*" + x3.Name
			}
		}

	case *ast.StarExpr:
		switch x2 := x.X.(type) {
		case *ast.Ident: // e.g. *Profile
			fieldType = "*" + x2.Name

		case *ast.SelectorExpr: // e.g. *time.Time
			switch x3 := x2.X.(type) {
			case *ast.Ident:
				fieldType = "*" + x3.Name + "." + x2.Sel.Name
			}
		}
	}

	var jsonName string
	var fieldTag string

	if field.Tag == nil {
		jsonName = fieldName
		fieldTag = ""
	} else {
		s, err := strconv.Unquote(field.Tag.Value)
		if err != nil {
			s = field.Tag.Value
		}

		jsonName = strings.Split((reflect.StructTag)(s).Get("json"), ",")[0]
		fieldTag = field.Tag.Value
	}

	fs := Field{
		Name:     fieldName,
		JSONName: jsonName,
		Type:     fieldType,
		Tag:      fieldTag,
	}

	return &fs, nil
}

func ParseModel(path string, exclude ...string) (Models, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}
	excludedModels := strings.Join(exclude, ",")

	models := []*Model{}

	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.GenDecl:
			if x.Tok != token.TYPE {
				break
			}

			for _, spec := range x.Specs {
				var fields []*Field

				var modelName string

				switch x2 := spec.(type) {
				case *ast.TypeSpec:
					modelName = x2.Name.Name

					// only accept struct
					if _, ok := x2.Type.(*ast.StructType); !ok {
						continue
					}

					// ignore excluded model
					if strings.Contains(excludedModels, modelName) {
						continue
					}

					switch x3 := x2.Type.(type) {
					case *ast.StructType:
						for _, field := range x3.Fields.List {
							fs, err := parseField(field)
							if err != nil {
								return false
							}

							fields = append(fields, fs)
						}
					}

					models = append(models, &Model{
						Name:   modelName,
						Fields: fields,
					})
				}
			}
		}

		return true
	})

	return models, nil
}
