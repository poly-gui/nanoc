package swiftgen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"nanoc/internal/datatype"
	"nanoc/internal/generator"
	"nanoc/internal/npschema"
)

type stringGenerator struct{}

func (g stringGenerator) TypeDeclaration(dataType datatype.DataType) string {
	return "String"
}

func (g stringGenerator) ReadSizeExpression(dataType datatype.DataType, varName string) string {
	if dataType.Kind == datatype.Enum {
		return fmt.Sprintf("%v.rawValue.lengthOfBytes(using: .utf8)", varName)
	}
	return fmt.Sprintf("%v.lengthOfBytes(using: .utf8)", varName)
}

func (g stringGenerator) ConstructorFieldParameter(field npschema.MessageField) string {
	return fmt.Sprintf("%v: %v", strcase.ToLowerCamel(field.Name), g.TypeDeclaration(field.Type))
}

func (g stringGenerator) FieldInitializer(field npschema.MessageField) string {
	c := strcase.ToLowerCamel(field.Name)
	return fmt.Sprintf("self.%v = %v", c, c)
}

func (g stringGenerator) FieldDeclaration(field npschema.MessageField) string {
	return fmt.Sprintf("let %v: %v", strcase.ToLowerCamel(field.Name), g.TypeDeclaration(field.Type))
}

func (g stringGenerator) ReadFieldFromBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	c := strcase.ToLowerCamel(field.Name)

	if field.Type.Kind == datatype.Enum {
		return generator.Lines(
			fmt.Sprintf("let %vSize = data.readSize(ofField: %d)", c, field.Number),
			fmt.Sprintf("guard let %vRawValue = data.read(at: ptr, withLength: %vSize) else {", c, c),
			"    return nil",
			"}",
			fmt.Sprintf("ptr += %vSize", c))
	}

	var v string
	var l4 string
	if ctx.IsVariableInScope(c) {
		v = c + "_"
		l4 = fmt.Sprintf("%v = %v", c, v)
	} else {
		v = c
	}

	return generator.Lines(
		fmt.Sprintf("let %vSize = data.readSize(ofField: %d)", c, field.Number),
		fmt.Sprintf("guard let %v = data.read(at: ptr, withLength: %vSize) else {", v, c),
		"    return nil",
		"}",
		l4,
		fmt.Sprintf("ptr += %vSize", c))
}

func (g stringGenerator) ReadValueFromBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	if dataType.Kind == datatype.Enum {
		return generator.Lines(
			fmt.Sprintf("let %vSize = data.readSize(at: ptr)", varName),
			"ptr += 4",
			fmt.Sprintf("guard let %vRawValue = data.read(at: ptr, withLength: %vSize) else {", varName, varName),
			"    return nil",
			"}",
			fmt.Sprintf("ptr += %vSize", varName))
	}

	var v string
	var l5 string
	if ctx.IsVariableInScope(varName) {
		v = varName + "_"
		l5 = fmt.Sprintf("%v = %v", varName, v)
	} else {
		v = varName
	}

	return generator.Lines(
		fmt.Sprintf("let %vSize = data.readSize(at: ptr)", varName),
		"ptr += 4",
		fmt.Sprintf("guard let %v = data.read(at: ptr, withLength: %vSize) else {", v, varName),
		"    return nil",
		"}",
		l5,
		fmt.Sprintf("ptr += %vSize", varName))
}

func (g stringGenerator) WriteFieldToBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	c := strcase.ToLowerCamel(field.Name)

	var expr string
	if field.Type.Kind == datatype.Enum {
		expr = c + ".rawValue"
	} else {
		expr = c
	}

	return generator.Lines(
		fmt.Sprintf("data.write(size: %v, ofField: %d, offset: offset)", g.ReadSizeExpression(field.Type, c), field.Number),
		fmt.Sprintf("data.append(string: %v)", expr))
}

func (g stringGenerator) WriteVariableToBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	var expr string
	if dataType.Kind == datatype.Enum {
		expr = varName + ".rawValue"
	} else {
		expr = varName
	}

	return generator.Lines(
		fmt.Sprintf("data.append(size: %v)", g.ReadSizeExpression(dataType, varName)),
		fmt.Sprintf("data.append(string: %v)", expr))
}
