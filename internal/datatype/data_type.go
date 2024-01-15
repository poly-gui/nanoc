package datatype

import (
	"nanoc/internal/symbol"
)

type DataType struct {
	// The name of this data type used in NanoPack schema.
	Identifier string

	// What kind of data type this is.
	Kind Kind

	ByteSize int

	// Schema is the schema for this message type.
	// Only applies to Enum or Message kind.
	Schema Schema

	// KeyType is the type of the keys stored in this data type.
	// Only applies to Map kind.
	KeyType *DataType

	// ElemType is the type of the element stored in this data type.
	// Only applies to Optional, Enum, Array and Map kind.
	ElemType *DataType
}

const DynamicSize = -1

type Kind uint

const (
	Int8 Kind = iota
	Int32
	Int64
	Double
	Bool
	String
	Array
	Map
	Enum
	Message
	Optional
)

var (
	npint8 = DataType{
		Identifier: "int8",
		Kind:       Int8,
		ByteSize:   1,
	}

	npint32 = DataType{
		Identifier: "int32",
		Kind:       Int32,
		ByteSize:   4,
	}

	npint64 = DataType{
		Identifier: "int64",
		Kind:       Int64,
		ByteSize:   8,
	}

	npdouble = DataType{
		Identifier: "double",
		Kind:       Double,
		ByteSize:   8,
	}

	npbool = DataType{
		Identifier: "bool",
		Kind:       Bool,
		ByteSize:   1,
	}

	npstring = DataType{
		Identifier: "string",
		Kind:       String,
		ByteSize:   DynamicSize,
	}
)

// SchemaMap is a map that maps names of schemas to the corresponding Schema definition.
type SchemaMap map[string]Schema

// FromKind returns the correct instance of DataType from the given Kind.
// Returns nil for non-string or non-primitive types.
func FromKind(kind Kind) *DataType {
	switch kind {
	case Int8:
		return &npint8
	case Int32:
		return &npint32
	case Int64:
		return &npint64
	case Double:
		return &npdouble
	case Bool:
		return &npdouble
	case String:
		return &npstring
	default:
		return nil
	}
}

// FromIdentifier returns the correct instance of DataType from the given identifier.
// Returns nil for non-string or non-primitive types.
func FromIdentifier(identifier string) *DataType {
	switch identifier {
	case npint8.Identifier:
		return &npint8

	case npint32.Identifier:
		return &npint32

	case npint64.Identifier:
		return &npint64

	case npdouble.Identifier:
		return &npdouble

	case npbool.Identifier:
		return &npbool

	case npstring.Identifier:
		return &npstring

	default:
		return nil
	}
}

func NewOptionalType(elemType *DataType) DataType {
	return DataType{
		Identifier: elemType.Identifier + symbol.Optional,
		Kind:       Optional,
		ByteSize:   DynamicSize,
		Schema:     nil,
		KeyType:    nil,
		ElemType:   elemType,
	}
}

func NewArrayType(elemType *DataType) DataType {
	return DataType{
		Identifier: elemType.Identifier + symbol.Array,
		Kind:       Array,
		ByteSize:   DynamicSize,
		Schema:     nil,
		KeyType:    nil,
		ElemType:   elemType,
	}
}

func NewMapType(keyType *DataType, valueType *DataType) DataType {
	return DataType{
		Identifier: symbol.MapBracketStart + keyType.Identifier + symbol.MapKeyValTypeSep + " " + valueType.Identifier + symbol.MapBracketEnd,
		Kind:       Map,
		ByteSize:   DynamicSize,
		Schema:     nil,
		KeyType:    keyType,
		ElemType:   valueType,
	}
}
