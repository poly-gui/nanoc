// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import Foundation
import NanoPack

func makeNanoPackMessage(from data: Data) -> NanoPackMessage? {
    let typeID = data.readTypeID()
    switch typeID {
    case 1676374721: return Widget(data: data)
    case 3495336243: return Text(data: data)
    default: return nil
    }
}
