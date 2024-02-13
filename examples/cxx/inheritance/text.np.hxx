// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef TEXT_NP_HXX
#define TEXT_NP_HXX

#include <nanopack/reader.hxx>
#include <string>
#include <vector>

#include "make_widget.np.hxx"
#include "widget.np.hxx"

struct Text : Widget {
  static constexpr int32_t TYPE_ID = 2;

  std::string content;

  Text() = default;

  Text(int32_t id, std::string content);

  Text(std::vector<uint8_t>::const_iterator begin, int &bytes_read);

  Text(const NanoPack::Reader &reader, int &bytes_read);

  [[nodiscard]] int32_t type_id() const override;

  [[nodiscard]] std::vector<uint8_t> data() const override;
};

#endif
