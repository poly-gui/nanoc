// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import Foundation
import NanoPack

let Person_typeID: TypeID = 1_225_883_824

class Person: NanoPackMessage {
  var typeID: TypeID { return 1_225_883_824 }

  var headerSize: Int { return 24 }

  let firstName: String
  let middleName: String?
  let lastName: String
  let age: Int8
  let otherFriend: Person?

  init(firstName: String, middleName: String?, lastName: String, age: Int8, otherFriend: Person?) {
    self.firstName = firstName
    self.middleName = middleName
    self.lastName = lastName
    self.age = age
    self.otherFriend = otherFriend
  }

  required init?(data: Data) {
    var ptr = data.startIndex + 24

    let firstNameSize = data.readSize(ofField: 0)
    guard let firstName = data.read(at: ptr, withLength: firstNameSize) else {
      return nil
    }
    ptr += firstNameSize

    var middleName: String?
    if data.readSize(ofField: 1) < 0 {
      middleName = nil
    } else {
      let middleNameSize = data.readSize(ofField: 1)
      guard let middleName_ = data.read(at: ptr, withLength: middleNameSize) else {
        return nil
      }
      middleName = middleName_
      ptr += middleNameSize
    }

    let lastNameSize = data.readSize(ofField: 2)
    guard let lastName = data.read(at: ptr, withLength: lastNameSize) else {
      return nil
    }
    ptr += lastNameSize

    let age: Int8 = data.read(at: ptr)
    ptr += 1

    var otherFriend: Person?
    if data.readSize(ofField: 4) < 0 {
      otherFriend = nil
    } else {
      let otherFriendByteSize = data.readSize(ofField: 4)
      guard let otherFriend_ = Person(data: data[ptr...]) else {
        return nil
      }
      otherFriend = otherFriend_
      ptr += otherFriendByteSize
    }

    self.firstName = firstName
    self.middleName = middleName
    self.lastName = lastName
    self.age = age
    self.otherFriend = otherFriend
  }

  required init?(data: Data, bytesRead: inout Int) {
    var ptr = data.startIndex + 24

    let firstNameSize = data.readSize(ofField: 0)
    guard let firstName = data.read(at: ptr, withLength: firstNameSize) else {
      return nil
    }
    ptr += firstNameSize

    var middleName: String?
    if data.readSize(ofField: 1) < 0 {
      middleName = nil
    } else {
      let middleNameSize = data.readSize(ofField: 1)
      guard let middleName_ = data.read(at: ptr, withLength: middleNameSize) else {
        return nil
      }
      middleName = middleName_
      ptr += middleNameSize
    }

    let lastNameSize = data.readSize(ofField: 2)
    guard let lastName = data.read(at: ptr, withLength: lastNameSize) else {
      return nil
    }
    ptr += lastNameSize

    let age: Int8 = data.read(at: ptr)
    ptr += 1

    var otherFriend: Person?
    if data.readSize(ofField: 4) < 0 {
      otherFriend = nil
    } else {
      let otherFriendByteSize = data.readSize(ofField: 4)
      guard let otherFriend_ = Person(data: data[ptr...]) else {
        return nil
      }
      otherFriend = otherFriend_
      ptr += otherFriendByteSize
    }

    self.firstName = firstName
    self.middleName = middleName
    self.lastName = lastName
    self.age = age
    self.otherFriend = otherFriend

    bytesRead = ptr - data.startIndex
  }

  func write(to data: inout Data, offset: Int) -> Int {
    let dataCountBefore = data.count

    data.reserveCapacity(offset + 24)

    data.append(typeID: TypeID(Person_typeID))
    data.append([0], count: 5 * 4)

    data.write(size: firstName.lengthOfBytes(using: .utf8), ofField: 0, offset: offset)
    data.append(string: firstName)

    if let middleName = self.middleName {
      data.write(size: middleName.lengthOfBytes(using: .utf8), ofField: 1, offset: offset)
      data.append(string: middleName)
    } else {
      data.write(size: -1, ofField: 1, offset: offset)
    }

    data.write(size: lastName.lengthOfBytes(using: .utf8), ofField: 2, offset: offset)
    data.append(string: lastName)

    data.write(size: 1, ofField: 3, offset: offset)
    data.append(int: age)

    if let otherFriend = self.otherFriend {
      let otherFriendByteSize = otherFriend.write(to: &data, offset: data.count)
      data.write(size: otherFriendByteSize, ofField: 4, offset: offset)
    } else {
      data.write(size: -1, ofField: 4, offset: offset)
    }

    return data.count - dataCountBefore
  }

  func data() -> Data? {
    var data = Data()
    _ = write(to: &data, offset: 0)
    return data
  }
}
