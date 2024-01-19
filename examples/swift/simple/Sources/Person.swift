// AUTOMATICALLY GENERATED BY NANOPACK. DO NOT MODIFY BY HAND.

import Foundation
import NanoPack

let Person_typeID: TypeID = 1

class Person: NanoPackMessage {
  var typeID: TypeID { return 1 }

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
    var ptr = 24

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
    var ptr = 24

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

    bytesRead = ptr
  }

  func data() -> Data? {
    var data = Data()
    data.reserveCapacity(24)

    withUnsafeBytes(of: Int32(Person_typeID)) {
      data.append(contentsOf: $0)
    }

    data.append([0], count: 5 * 4)

    data.write(size: firstName.lengthOfBytes(using: .utf8), ofField: 0)
    data.append(string: firstName)

    if let middleName = self.middleName {
      data.write(size: middleName.lengthOfBytes(using: .utf8), ofField: 1)
      data.append(string: middleName)
    } else {
      data.write(size: -1, ofField: 1)
    }

    data.write(size: lastName.lengthOfBytes(using: .utf8), ofField: 2)
    data.append(string: lastName)

    data.write(size: 1, ofField: 3)
    data.append(int: age)

    if let otherFriend = self.otherFriend {
      guard let otherFriendData = otherFriend.data() else {
        return nil
      }
      data.write(size: otherFriendData.count, ofField: 4)
      data.append(otherFriendData)
    } else {
      data.write(size: -1, ofField: 4)
    }

    return data
  }
}