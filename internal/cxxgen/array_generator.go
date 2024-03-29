package cxxgen

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
	ig := g.gm[dataType.ElemType.Kind]
	return fmt.Sprintf("std::vector<%v>", ig.TypeDeclaration(*dataType.ElemType))
}

func (g arrayGenerator) ReadSizeExpression(dataType datatype.DataType, varName string) string {
	if dataType.ElemType.ByteSize != datatype.DynamicSize {
		return fmt.Sprintf("%v.size() * %d", varName, dataType.ElemType.ByteSize)
	}
	return fmt.Sprintf("%v_byte_size", varName)
}

func (g arrayGenerator) ConstructorFieldParameter(field npschema.MessageField) string {
	return g.TypeDeclaration(field.Type) + " " + strcase.ToSnake(field.Name)
}

func (g arrayGenerator) FieldInitializer(field npschema.MessageField) string {
	s := strcase.ToSnake(field.Name)
	return fmt.Sprintf("%v(std::move(%v))", s, s)
}

func (g arrayGenerator) FieldDeclaration(field npschema.MessageField) string {
	return g.TypeDeclaration(field.Type) + " " + strcase.ToSnake(field.Name) + ";"
}

func (g arrayGenerator) ReadFieldFromBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	s := strcase.ToSnake(field.Name)
	var l1 string
	if field.Type.ElemType.ByteSize != datatype.DynamicSize {
		// for arrays with fixed size items, the number of elements in the array can be calculated.
		l1 = generator.Lines(
			fmt.Sprintf("const int32_t %v_byte_size = reader.read_field_size(%d);", s, field.Number),
			fmt.Sprintf("const int32_t %v_vec_size = %v_byte_size / %d;", s, s, field.Type.ElemType.ByteSize))
		ctx.AddVariableToScope(s + "_vec_size")
	}
	return generator.Lines(
		l1,
		g.ReadValueFromBuffer(field.Type, s, ctx),
		fmt.Sprintf("this->%v = std::move(%v);", s, s))
}

func (g arrayGenerator) ReadValueFromBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	i32g := g.gm[datatype.Int32]
	ig := g.gm[dataType.ElemType.Kind]
	vecSizeVar := varName + "_vec_size"

	lv := ctx.NewLoopVar()

	// If the number of elements in the vector is not read previously,
	// generate code to read it here.
	var l0 string
	if !ctx.IsVariableInScope(vecSizeVar) {
		l0 = i32g.ReadValueFromBuffer(*datatype.FromKind(datatype.Int32), vecSizeVar, ctx)
	}

	var l1 string
	if ctx.IsVariableInScope(varName) {
		l1 = fmt.Sprintf("%v.reserve(%v);", varName, vecSizeVar)
	} else {
		l1 = generator.Lines(
			fmt.Sprintf("%v %v;", g.TypeDeclaration(dataType), varName),
			fmt.Sprintf("%v.reserve(%v);", varName, vecSizeVar))
	}

	var l4 string
	if isTriviallyCopiable(dataType) {
		l4 = fmt.Sprintf("%v.emplace_back(%v);", varName, lv+"_item")
	} else {
		l4 = fmt.Sprintf("%v.emplace_back(std::move(%v));", varName, lv+"_item")
	}

	ls := generator.Lines(
		l0,
		l1,
		fmt.Sprintf("for (int %v = 0; %v < %v; %v++) {", lv, lv, vecSizeVar, lv),
		ig.ReadValueFromBuffer(*dataType.ElemType, lv+"_item", ctx),
		l4,
		"}")

	ctx.RemoveVariableFromScope(lv)

	return ls
}

func (g arrayGenerator) WriteFieldToBuffer(field npschema.MessageField, ctx generator.CodeContext) string {
	s := strcase.ToSnake(field.Name)

	if field.Type.ElemType.ByteSize != datatype.DynamicSize {
		// the array has fixed size elements, so the total size of the array
		// data can be calculated directly:
		//
		//     sizeof(element type) * number of elements in the vector +
		//
		// the number of elements in the array is not written to the buffer,
		// unlike elements with dynamic size, since it can be determined easily:
		//
		//     number of elements in the vector = total byte size of vector /
		//     sizeof(element type).
		lv := ctx.NewLoopVar()
		ig := g.gm[field.Type.ElemType.Kind]
		ls := generator.Lines(
			fmt.Sprintf("const int32_t %v_byte_size = %v.size() * %d;", s, s, field.Type.ElemType.ByteSize),
			fmt.Sprintf("NanoPack::write_field_size(%d, %v_byte_size, offset, buf);", field.Number, s),
			fmt.Sprintf("for (const auto &%v : %v) {", lv, s),
			ig.WriteVariableToBuffer(*field.Type.ElemType, lv, ctx),
			"}",
			fmt.Sprintf("bytes_written += %v_byte_size;", s))
		ctx.RemoveVariableFromScope(lv)
		return ls
	}

	return generator.Lines(
		g.WriteVariableToBuffer(field.Type, s, ctx),
		fmt.Sprintf("NanoPack::write_field_size(%d, %v_byte_size, offset, buf);", field.Number, s),
		fmt.Sprintf("bytes_written += %v_byte_size;", s))
}

func (g arrayGenerator) WriteVariableToBuffer(dataType datatype.DataType, varName string, ctx generator.CodeContext) string {
	ig := g.gm[dataType.ElemType.Kind]
	i32g := g.gm[datatype.Int32]
	vecSizeVar := varName + "_vec_size"
	isItemDynamicSize := dataType.ElemType.ByteSize == datatype.DynamicSize

	l0 := fmt.Sprintf("const size_t %v = %v.size();", vecSizeVar, varName)
	l1 := i32g.WriteVariableToBuffer(*datatype.FromKind(datatype.Int32), vecSizeVar, ctx)

	var l2 string
	if isItemDynamicSize {
		// elements in the array are dynamically sized,
		// so the total byte size of the array cannot be determined statically.
		// here we declare a variable for storing the total byte size of all the
		// elements in the array, which is accumulated later in the loop
		l2 = fmt.Sprintf("int32_t %v_byte_size = sizeof(int32_t);", varName)
	}

	lv := ctx.NewLoopVar()
	ls := generator.Lines(
		l0,
		l1,
		l2,
		fmt.Sprintf("for (auto &%v : %v) {", lv, varName),
		ig.WriteVariableToBuffer(*dataType.ElemType, lv, ctx))

	var l5 string
	if isItemDynamicSize {
		l5 = fmt.Sprintf("%v_byte_size += %v;", varName, ig.ReadSizeExpression(*dataType.ElemType, lv))
	}

	return generator.Lines(
		ls,
		l5,
		"}")
}
