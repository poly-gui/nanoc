// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import { NanoBufReader, type NanoPackMessage } from "nanopack";

import { Widget } from "./widget.np.js";
import { Text } from "./text.np.js";

function makeNanoPackMessage(
  bytes: Uint8Array,
): { bytesRead: number; result: NanoPackMessage } | null {
  const reader = new NanoBufReader(bytes);
  switch (reader.readTypeId()) {
    case 1676374721:
      return Widget.fromReader(reader);
    case 3495336243:
      return Text.fromReader(reader);
    default:
      return null;
  }
}

export { makeNanoPackMessage };
