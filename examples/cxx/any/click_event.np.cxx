// AUTOMATICALLY GENERATED BY NANOC

#include <nanopack/reader.hxx>
#include <nanopack/writer.hxx>

#include "click_event.np.hxx"

ClickEvent::ClickEvent(double x, double y, int64_t timestamp)
    : x(x), y(y), timestamp(timestamp) {}

size_t ClickEvent::read_from(NanoPack::Reader &reader) {
  uint8_t *buf = reader.buffer;
  int ptr = 16;

  reader.read_double(ptr, x);
  ptr += 8;

  reader.read_double(ptr, y);
  ptr += 8;

  reader.read_int64(ptr, timestamp);
  ptr += 8;

  return ptr;
}

NanoPack::TypeId ClickEvent::type_id() const { return TYPE_ID; }

size_t ClickEvent::header_size() const { return 16; }

size_t ClickEvent::write_to(NanoPack::Writer &writer, int offset) const {
  const size_t writer_size_before = writer.size();

  writer.reserve_header(16);

  writer.write_type_id(TYPE_ID, offset);

  writer.write_field_size(0, 8, offset);
  writer.append_double(x);

  writer.write_field_size(1, 8, offset);
  writer.append_double(y);

  writer.write_field_size(2, 8, offset);
  writer.append_int64(timestamp);

  return writer.size() - writer_size_before;
}
