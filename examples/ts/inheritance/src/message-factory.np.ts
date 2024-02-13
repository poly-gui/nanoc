// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import { NanoBufReader, type NanoPackMessage } from "nanopack";

import { Text } from "./text.np.js";
import { Widget } from "./widget.np.js";

function makeNanoPackMessage(
  bytes: Uint8Array,
): { bytesRead: number; result: NanoPackMessage } | null {
  const reader = new NanoBufReader(bytes);
  switch (reader.readTypeId()) {
    case 2:
      return Text.fromReader(reader);
    case 1:
      return Widget.fromReader(reader);
    default:
      return null;
  }
}

export { makeNanoPackMessage };
