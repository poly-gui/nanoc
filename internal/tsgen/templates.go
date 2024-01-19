package tsgen

import "nanoc/internal/npschema"

type messageClassTemplateInfo struct {
	Schema          *npschema.Message
	ExternalImports []string

	ConstructorParameters []string
	ConstructorArgs       []string
	SuperConstructorArgs  []string

	ReadPtrStart           int
	FieldReadCodeFragments []string

	InitialWriteBufferSize  int
	FieldWriteCodeFragments []string
}

type messageClassFactoryTemplateInfo struct {
	Schema             *npschema.Message
	MessageClassImport string
	MessageImports     []string
}

type messageFactoryTemplateInfo struct {
	Schemas        []*npschema.Message
	MessageImports []string
}

const (
	extImport = ".np.js"

	extTsFile = ".np.ts"
)

const (
	templateNameMessageClass = "TsMessageClass"

	templateNameMessageClassFactory = "TsMessageClassFactory"

	templateNameMessageFactory = "TsMessageFactory"
)

const fileNameMessageFactoryFile = "make-nanopack-message"

const messageClassTemplate = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import { NanoBufReader, NanoBufWriter{{if not .Schema.HasParentMessage}}, type NanoPackMessage{{end}} } from "nanopack";

{{range .ExternalImports}}
{{.}}
{{- end}}

class {{.Schema.Name}} {{if .Schema.HasParentMessage}}extends {{.Schema.ParentMessage.Name}}{{else}}implements NanoPackMessage{{end}} {
  public static TYPE_ID = {{.Schema.TypeID}};

  {{if .Schema.HasParentMessage -}}
  constructor({{join .ConstructorParameters ", "}}) {
    super({{join .SuperConstructorArgs ", "}})
  }
  {{else}}
  constructor({{join .ConstructorParameters ", "}}) {}
  {{- end}}

  public static fromBytes(bytes: Uint8Array): { bytesRead: number, result: {{.Schema.Name}} } | null {
    const reader = new NanoBufReader(bytes);
    return {{.Schema.Name}}.fromReader(reader);
  }

  public static fromReader(reader: NanoBufReader): { bytesRead: number, result: {{.Schema.Name}} } | null {
    let ptr = {{.ReadPtrStart}};

    {{range .FieldReadCodeFragments}}
    {{.}}

    {{end}}

    return { bytesRead: ptr, result: new {{.Schema.Name}}({{join .ConstructorArgs ", "}}) };
  }

  {{if .Schema.HasParentMessage}}override {{end}}public get typeId(): number {
    return {{.Schema.TypeID}}; 
  }

  {{if .Schema.HasParentMessage}}override {{end}}public bytes(): Uint8Array {
    const writer = new NanoBufWriter({{.InitialWriteBufferSize}});
    writer.writeTypeId({{.Schema.TypeID}});

    {{range .FieldWriteCodeFragments}}
    {{.}}

    {{end}}

    return writer.bytes;
  }

  {{if .Schema.HasParentMessage}}override {{end}}public bytesWithLengthPrefix(): Uint8Array {
    const writer = new NanoBufWriter({{.InitialWriteBufferSize}} + 4, true);
    writer.writeTypeId({{.Schema.TypeID}});

    {{range .FieldWriteCodeFragments}}
    {{.}}

    {{end}}

    writer.writeLengthPrefix(writer.currentSize - 4);

    return writer.bytes;
  }
}

export { {{.Schema.Name}} };
`

const messageClassFactoryTemplate = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import { NanoBufReader } from "nanopack";

import { {{.Schema.Name}} } from "{{.MessageClassImport}}";
{{range .MessageImports -}}
{{.}}
{{- end}}

function make{{.Schema.Name}}(bytes: Uint8Array) {
  const reader = new NanoBufReader(bytes);
  switch (reader.readTypeId()) {
  case {{.Schema.TypeID}}: return {{.Schema.Name}}.fromReader(reader);
  {{range .Schema.ChildMessages}}
  case {{.TypeID}}: return {{.Name}}.fromReader(reader);
  {{- end}}
  default: return null;
  }
}

export { make{{.Schema.Name}} } ;
`

const messageFactoryTemplate = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import type { NanoPackMessage } from "nanopack";
{{range .MessageImports}}
{{.}}
{{- end}}

function makeNanoPackMessage(bytes: Uint8Array, typeId: number): { bytesRead: number, result: NanoPackMessage } | null {
  switch (typeId) {
  {{range .Schemas}}
  case {{.TypeID}}: return {{.Name}}.fromBytes(bytes);
  {{- end}}
  default: return null;
  }
}

export { makeNanoPackMessage }
`
