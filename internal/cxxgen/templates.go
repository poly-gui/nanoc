package cxxgen

import "nanoc/internal/npschema"

type messageHeaderFileTemplateInfo struct {
	Namespace         string
	MessageName       string
	HasParentMessage  bool
	ParentMessageName string
	TypeID            int
	IncludeGuardName  string
	LibraryImports    []string
	RelativeImports   []string
	IsInherited       bool

	FieldDeclarationLines []string
	ConstructorParameters []string
}

type messageImplFileTemplateInfo struct {
	Namespace         string
	HeaderName        string
	MessageName       string
	HasParentMessage  bool
	ParentMessageName string

	ConstructorParameters []string
	SuperConstructorArgs  []string
	FieldInitializers     []string

	ReadPtrStart           int
	FieldReadCodeFragments []string

	InitialWriteBufferSize  int
	FieldWriteCodeFragments []string
}

type childMessageFactoryHeaderFileTemplateInfo struct {
	Namespace           string
	IncludeGuardName    string
	MessageName         string
	MessageHeaderName   string
	FactoryFunctionName string
}

type childMessageFactoryImplFileTemplateInfo struct {
	Namespace               string
	Schema                  *npschema.Message
	ChildMessageImportPaths []string
	HeaderName              string
	FactoryFunctionName     string
}

type messageFactoryHeaderFileTemplateInfo struct {
	Namespace string
}

type messageFactoryImplFileTemplateInfo struct {
	Namespace          string
	MessageImportPaths []string
	MessageSchemas     []*npschema.Message
}

type enumHeaderFileInfo struct {
	Namespace        string
	Schema           *npschema.Enum
	BackingTypeName  string
	MemberNames      []string
	IncludeGuardName string
}

const (
	templateNameMessageHeaderFile             = "CxxMessageHeaderFile"
	templateNameMessageImplFile               = "CxxMessageImplFile"
	templateNameMessageFactoryHeaderFile      = "CxxMessageFactoryHeaderFile"
	templateNameMessageFactoryImplFile        = "CxxMessageFactoryImplFile"
	templateNameChildMessageFactoryHeaderFile = "CxxChildMessageFactoryHeaderFile"
	templateNameChildMessageFactoryImplFile   = "CxxChildMessageFactoryImplFile"
	templateNameEnumHeaderFile                = "CxxEnumHeaderFile"
)

const (
	extHeaderFile = ".np.hxx"
	extImplFile   = ".np.cxx"
)

const cxxSymbolMemberOf = "::"

const fileNameMessageFactory = "nanopack_message_factory"

const messageHeaderFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef {{.IncludeGuardName}}
#define {{.IncludeGuardName}}

#include <vector>
{{- range .LibraryImports}}
#include <{{.}}>
{{- end}}
{{- if not .HasParentMessage}}
#include <nanopack/message.hxx>
{{- end}}
#include <nanopack/reader.hxx>

{{range .RelativeImports}}
#include "{{.}}"
{{- end}}

{{if .Namespace}}namespace {{.Namespace}} { {{- end}}

struct {{.MessageName}} : {{if .HasParentMessage}}{{.ParentMessageName}}{{else}}NanoPack::Message{{end}} {
  static constexpr int32_t TYPE_ID = {{.TypeID}};

  {{range .FieldDeclarationLines}}{{.}}{{end}}

  {{.MessageName}}() = default;

  {{$l := len .ConstructorParameters}}{{if eq $l 1}}explicit {{end}}
  {{- .MessageName}}({{range $i, $v := .ConstructorParameters}}{{if $i}}, {{end}}{{$v}}{{end}});

  {{.MessageName}}(std::vector<uint8_t>::const_iterator begin, int &bytes_read);

  {{.MessageName}}(const NanoPack::Reader &reader, int &bytes_read);

  [[nodiscard]] int32_t type_id() const override;

  [[nodiscard]] std::vector<uint8_t> data() const override;
};

{{if .Namespace}}} // namespace {{.Namespace}}{{end}}

#endif
`
const messageImplFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#include <nanopack/reader.hxx>
#include <nanopack/writer.hxx>

#include "{{.HeaderName}}"

{{if .Namespace}}{{.Namespace}}::{{end}}{{.MessageName}}::{{.MessageName}}({{join .ConstructorParameters ", "}}) :
  {{if .HasParentMessage}}{{.ParentMessageName}}({{join .SuperConstructorArgs ", "}}), {{end}}
  {{- join .FieldInitializers ", "}} {}

{{if .Namespace}}{{.Namespace}}::{{end}}{{.MessageName}}::{{.MessageName}}(const NanoPack::Reader &reader, int &bytes_read) {{if .HasParentMessage}}: {{.ParentMessageName}}(){{end}} {
  const auto begin = reader.begin();
  int ptr = {{.ReadPtrStart}};

  {{range .FieldReadCodeFragments}}
  {{.}}

  {{end}}

  bytes_read = ptr;
}

{{if .Namespace}}{{.Namespace}}::{{end}}{{.MessageName}}::{{.MessageName}}(std::vector<uint8_t>::const_iterator begin, int &bytes_read) :
  {{.MessageName}}(NanoPack::Reader(begin), bytes_read) {}

int32_t {{if .Namespace}}{{.Namespace}}::{{end}}{{.MessageName}}::type_id() const {
  return TYPE_ID;
}

std::vector<uint8_t> {{if .Namespace}}{{.Namespace}}::{{end}}{{.MessageName}}::data() const {
  std::vector<uint8_t> buf({{.InitialWriteBufferSize}});
  NanoPack::Writer writer(&buf);

  writer.write_type_id(TYPE_ID);

  {{range .FieldWriteCodeFragments}}
  {{.}}

  {{end}}

  return buf;
}
`

const childMessageFactoryHeaderFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef {{.IncludeGuardName}}
#define {{.IncludeGuardName}}

#include <memory>
#include "{{.MessageHeaderName}}"

{{if .Namespace}}namespace {{.Namespace}} { {{- end}}

std::unique_ptr<{{.MessageName}}> {{.FactoryFunctionName}}(std::vector<uint8_t>::const_iterator begin, int &bytes_read);

{{if .Namespace}}} // namespace {{.Namespace}}{{end}}

#endif
`

const childMessageFactoryImplFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#include <nanopack/reader.hxx>
#include "{{.HeaderName}}"
{{range .ChildMessageImportPaths}}
#include "{{.}}"
{{- end}}

std::unique_ptr<{{.Schema.Name}}> {{if .Namespace}}{{.Namespace}}::{{end}}{{.FactoryFunctionName}}(std::vector<uint8_t>::const_iterator begin, int &bytes_read) {
  const NanoPack::Reader reader(begin);
  switch (reader.read_type_id()) {
  case {{.Schema.TypeID}}: return std::make_unique<{{.Schema.Name}}>(reader, bytes_read);
  {{range .Schema.ChildMessages -}}
  case {{.TypeID}}: return std::make_unique<{{.Name}}>(reader, bytes_read);
  {{end -}}
  default: return nullptr;
  }
}
`

const messageFactoryHeaderFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef NANOPACK_MESSAGE_FACTORY_HXX
#define NANOPACK_MESSAGE_FACTORY_HXX

#include <nanopack/message.hxx>
#include <memory>

{{if .Namespace}}namespace {{.Namespace}} { {{- end}}

std::unique_ptr<NanoPack::Message> make_nanopack_message(std::vector<uint8_t>::const_iterator data_iter, int &bytes_read);

{{if .Namespace}}} // namespace {{.Namespace}}{{end}}

#endif
`

const messageFactoryImplFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#include "nanopack_message_factory.np.hxx"

{{range .MessageImportPaths}}
#include "{{.}}"
{{- end}}

std::unique_ptr<NanoPack::Message> {{if .Namespace}}{{.Namespace}}::{{end}}make_nanopack_message(std::vector<uint8_t>::const_iterator data_iter, int &bytes_read) {
  switch (const NanoPack::Reader reader(data_iter); reader.read_type_id()) {
  {{range .MessageSchemas -}}
  case {{.TypeID}}: return std::make_unique<{{.Name}}>(reader, bytes_read);
  {{- end}}
  default: return nullptr;
  }
}
`

const enumHeaderFile = `// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef {{.IncludeGuardName}}
#define {{.IncludeGuardName}}

#include <array>
#include <stdexcept>

{{if .Namespace}}namespace {{.Namespace}} { {{- end}}

class {{.Schema.Name}} {
public:
  enum {{.Schema.Name}}Member {
  	{{- range .MemberNames}}
    {{.}},
    {{- end}}
  };

private:
  constexpr static std::array<{{.BackingTypeName}}, {{len .MemberNames}}> values = { {{range .Schema.Members}}{{.ValueLiteral}}, {{end}} };
  {{- if eq .BackingTypeName "std::string_view" -}}
  inline static std::unordered_map<std::string_view, {{.Schema.Name}}Member> lookup{
    {{- $m := .MemberNames -}}
    {{range $i, $e := .Schema.Members}}{{if $i}}, {{end}}{ {{$e.ValueLiteral}}, {{index $m $i}} }{{end}}
  };
  {{- end -}}

  {{.Schema.Name}}Member enum_value;
  {{.BackingTypeName}} _value;

public:
  {{.Schema.Name}}() = default;

  {{if eq .BackingTypeName "std::string_view" -}}
  explicit {{.Schema.Name}}(const {{.BackingTypeName}} &value) : enum_value(lookup.find(value)->second), _value(values[enum_value]) {}
  {{- else -}}
  explicit {{.Schema.Name}}(const {{.BackingTypeName}} &value) {
    switch (value) {
      {{range .Schema.Members}}case {{.ValueLiteral}}:
        enum_value = {{.Name}};
        break;{{end}}
      default: throw std::runtime_error("invalid value for enum {{.Schema.Name}}");
    }
    _value = values[enum_value];
  }
  {{- end}}

  constexpr {{.Schema.Name}}({{.Schema.Name}}Member member) : enum_value(member), _value(values[member]) {}

  [[nodiscard]] constexpr const {{.BackingTypeName}} &value() const { return _value; }

  constexpr operator {{.Schema.Name}}Member() const { return enum_value; }

  explicit operator bool() const = delete;
};

{{if .Namespace}}} // namespace {{.Namespace}}{{end}}

#endif
`
