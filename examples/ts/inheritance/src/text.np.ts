// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import { NanoBufReader, NanoBufWriter } from "nanopack";

import { Widget } from "./widget.np.js";

class Text extends Widget {
  public static TYPE_ID = 3495336243;

  public override readonly typeId: number = 3495336243;

  public override readonly headerSize: number = 12;

  constructor(
    public id: number,
    public content: string,
  ) {
    super(id);
  }

  public static fromBytes(
    bytes: Uint8Array,
  ): { bytesRead: number; result: Text } | null {
    const reader = new NanoBufReader(bytes);
    return Text.fromReader(reader);
  }

  public static fromReader(
    reader: NanoBufReader,
  ): { bytesRead: number; result: Text } | null {
    let ptr = 12;

    const id = reader.readInt32(ptr);
    ptr += 4;

    const contentByteLength = reader.readFieldSize(1);
    const content = reader.readString(ptr, contentByteLength);
    ptr += contentByteLength;

    return { bytesRead: ptr, result: new Text(id, content) };
  }

  public override writeTo(writer: NanoBufWriter, offset: number = 0): number {
    let bytesWritten = 12;

    writer.writeTypeId(3495336243, offset);

    writer.appendInt32(this.id);
    writer.writeFieldSize(0, 4, offset);
    bytesWritten += 4;

    const contentByteLength = writer.appendString(this.content);
    writer.writeFieldSize(1, contentByteLength, offset);
    bytesWritten += contentByteLength;

    return bytesWritten;
  }

  public override bytes(): Uint8Array {
    const writer = new NanoBufWriter(12);
    this.writeTo(writer);
    return writer.bytes;
  }
}

export { Text };
