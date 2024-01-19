package tsgen

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"nanoc/internal/datatype"
	"nanoc/internal/generator"
	"nanoc/internal/npschema"
)

type arrayGenerator struct {
	gm generator.MessageCodeGeneratorMap
}

func (g arrayGenerator) TypeDeclaration(dataType datatype.DataType) string {
	ig := g.gm[dataType.Kind]
	return ig.TypeDeclaration(*dataType.ElemType) + "[]"
}

func (g arrayGenerator) ReadSizeExpression(dataType datatype.DataType, varName string) string {
	if dataType.ElemType.ByteSize == datatype.DynamicSize {
		return fmt.Sprintf("%vByteLength", varName)
	}
	return fmt.Sprintf("%v.length * %d", varName, dataType.ElemType.ByteSize)
}

func (g arrayGenerator) ConstructorFieldParameter(field npschema.MessageField) string {
	return fmt.Sprintf("public %v: %v", strcase.ToLowerCamel(field.Name), g.TypeDeclaration(field.Type))
}

func (g arrayGenerator) FieldInitializer(field npschema.MessageField) string {
	c := strcase.ToLowerCamel(field.Name)
	return fmt.Sprintf("this.%v = %v;", c, c)
}

func (g arrayGenerator) FieldDeclaration(field npschema.MessageField) string {
	return fmt.Sprintf("public %v: %v", strcase.ToLowerCamel(field.Name), g.TypeDeclaration(field.Type))
}

func (g arrayGenerator) ReadFieldFromBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	c := strcase.ToLowerCamel(field.Name)

	if field.Type.ElemType.ByteSize == datatype.DynamicSize {
		return g.ReadValueFromBuffer(field.Type, c, ctx)
	}

	ctx.AddVariableToScope(c + "ByteLength")
	ctx.AddVariableToScope(c + "ItemCount")

	return generator.Lines(
		fmt.Sprintf("const %vByteLength = reader.readFieldSize(%d);", c, field.Number),
		fmt.Sprintf("const %vLength = %vByteLength / %d;", c, c, field.Type.ElemType.ByteSize),
		g.ReadValueFromBuffer(field.Type, c, ctx))
}

func (g arrayGenerator) ReadValueFromBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	i32g := g.gm[datatype.Int32]
	ig := g.gm[dataType.ElemType.Kind]

	var l0 string
	if !ctx.IsVariableInScope(varName + "Length") {
		l0 = i32g.ReadValueFromBuffer(*datatype.FromKind(datatype.Int32), varName+"Length", ctx)
	}

	var l1 string
	if ctx.IsVariableInScope(varName) {
		l1 = fmt.Sprintf("%v = new Array(%vLength);", varName, varName)
	} else {
		l1 = fmt.Sprintf("const %v: %v = new Array(%vLength);", varName, g.TypeDeclaration(dataType), varName)
	}

	lv := ctx.NewLoopVar()
	iv := lv + "Item"
	ctx.AddVariableToScope(iv)

	ls := generator.Lines(
		l0,
		l1,
		fmt.Sprintf("for (let %v = 0; %v < %vLength; %v++) {", lv, lv, varName, lv),
		ig.ReadValueFromBuffer(*dataType.ElemType, iv, ctx),
		fmt.Sprintf("%v[%v] = %v", varName, lv, iv),
		"}")

	ctx.RemoveVariableFromScope(iv)
	ctx.RemoveVariableFromScope(lv)

	return ls
}

func (g arrayGenerator) WriteFieldToBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	c := strcase.ToLowerCamel(field.Name)
	ig := g.gm[field.Type.ElemType.Kind]

	if field.Type.ElemType.ByteSize == datatype.DynamicSize {
		return generator.Lines(
			g.WriteVariableToBuffer(field.Type, "this."+c, ctx),
			fmt.Sprintf("writer.writeFieldSize(%d, %vByteLength);", field.Number, c))
	}

	lv := ctx.NewLoopVarWithSuffix("Item")

	ls := generator.Lines(
		fmt.Sprintf("writer.writeFieldSize(%d, %v.length * %d);", field.Number, c, field.Type.ElemType.ByteSize),
		fmt.Sprintf("for (const %v of this.%v) {", c, c),
		ig.WriteVariableToBuffer(*field.Type.ElemType, c, ctx),
		"}")

	ctx.RemoveVariableFromScope(lv)

	return ls
}

func (g arrayGenerator) WriteVariableToBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	ig := g.gm[dataType.ElemType.Kind]

	lv := ctx.NewLoopVarWithSuffix("Item")

	var l1 string
	var l4 string
	if dataType.ElemType.ByteSize == datatype.DynamicSize {
		l1 = fmt.Sprintf("let %vByteLength = 4;", varName)
		l4 = fmt.Sprintf("%vByteLength += %v;", varName, ig.ReadSizeExpression(*dataType.ElemType, lv))
	}

	return generator.Lines(
		fmt.Sprintf("writer.appendInt32(%v.length);", varName),
		l1,
		fmt.Sprintf("for (const %v of %v) {", lv, varName),
		ig.WriteVariableToBuffer(*dataType.ElemType, lv, ctx),
		l4,
		"}",
	)
}
