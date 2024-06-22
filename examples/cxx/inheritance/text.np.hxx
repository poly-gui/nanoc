// AUTOMATICALLY GENERATED BY NANOC

#ifndef TEXT_NP_HXX
#define TEXT_NP_HXX
#include <nanopack/nanopack.hxx>
#include <nanopack/reader.hxx>
#include <string>

#include "make_widget.np.hxx"
#include "widget.np.hxx"

struct Text : Widget {
  static constexpr NanoPack::TypeId TYPE_ID = 3495336243;

  std::string content;

  Text() = default;

  Text(int32_t id, std::string content);

  size_t read_from(NanoPack::Reader &reader);

  size_t write_to(NanoPack::Writer &writer, int offset) const override;

  [[nodiscard]] NanoPack::TypeId type_id() const override;

  [[nodiscard]] size_t header_size() const override;
};

#endif
