package nanoc

import (
	"errors"
	"nanoc/internal/datatype"
	"nanoc/internal/npschema"
	"nanoc/internal/tsgen"
	"reflect"
)

func runTSSchemaGenerator(schema datatype.Schema, opts Options) error {
	switch s := schema.(type) {
	case *npschema.Message:
		return tsgen.GenerateMessageClass(s, tsgen.Options{
			FormatterPath:      opts.CodeFormatterPath,
			FormatterArgs:      opts.CodeFormatterArgs,
			MessageFactoryPath: opts.MessageFactoryAbsFilePath,
		})
	default:
		return errors.New("unexpected error. Unsupported schema type " + reflect.TypeOf(schema).Name())
	}
}

func runTSMessageFactoryGenerator(schemas []*npschema.Message, opts Options) error {
	return tsgen.GenerateMessageFactory(schemas, tsgen.Options{
		FormatterPath:      opts.CodeFormatterPath,
		FormatterArgs:      opts.CodeFormatterArgs,
		MessageFactoryPath: opts.MessageFactoryAbsFilePath,
	})
}
