// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef COLUMN_NP_HXX
#define COLUMN_NP_HXX

#include <nanopack/message.hxx>
#include <nanopack/nanopack.hxx>
#include <nanopack/reader.hxx>
#include <vector>

#include "alignment.np.hxx"

struct Column : NanoPack::Message {
  static constexpr NanoPack::TypeId TYPE_ID = 2415007766;

  Alignment alignment;

  Column() = default;

  explicit Column(const Alignment &alignment);

  Column(std::vector<uint8_t>::const_iterator begin, int &bytes_read);

  Column(const NanoPack::Reader &reader, int &bytes_read);

  size_t write_to(std::vector<uint8_t> &buf, int offset) const override;

  [[nodiscard]] NanoPack::TypeId type_id() const override;

  [[nodiscard]] int header_size() const override;

  [[nodiscard]] std::vector<uint8_t> data() const override;
};

#endif
