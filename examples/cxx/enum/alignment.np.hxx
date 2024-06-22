// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

#ifndef ALIGNMENT_ENUM_NP_HXX
#define ALIGNMENT_ENUM_NP_HXX

#include <array>
#include <string_view>
#include <unordered_map>

class Alignment {
public:
  enum AlignmentMember {
    START,
    END,
    CENTER,
    TOP,
    BOTTOM,
  };

private:
  constexpr static std::array<std::string_view, 5> values = {
      "start", "end", "center", "top", "bottom",
  };
  inline static std::unordered_map<std::string_view, AlignmentMember> lookup{
      {"start", START},
      {"end", END},
      {"center", CENTER},
      {"top", TOP},
      {"bottom", BOTTOM}};
  AlignmentMember enum_value;
  std::string_view _value;

public:
  Alignment() = default;

  explicit Alignment(const std::string_view &value)
      : enum_value(lookup.find(value)->second), _value(values[enum_value]) {}

  constexpr Alignment(AlignmentMember member)
      : enum_value(member), _value(values[member]) {}

  [[nodiscard]] constexpr const std::string_view &value() const {
    return _value;
  }

  constexpr operator AlignmentMember() const { return enum_value; }

  explicit operator bool() const = delete;
};

#endif
