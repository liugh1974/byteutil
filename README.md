## Encode/Decode binary data by go

## How to use it

### ByteEncoder/ByteDecoder

en := byteutil.NewByteEncoder()

en.WriteInt(1, 10)

en.WriteInt(4, 1000000)

en.WriteInt64(5, 10000000)

en.WriteString("example")


data := en.GetContent()

de := byteuit.NewByteDecoder(data)

val, _ := de.ReadInt(1)   // val is 10

val, _ = de.ReadInt(4)    // val is 1000000

num, _ := de.ReadInt64(5) // num is 10000000

str, _ := de.ReadString() // str is "example"


### BitEncoder/BitDecoder

bitEn := byteutil.NewBitEncoder()

bitEn.WriteInt(4, 10)

bitEn.WriteBool(true)

bitEn.WriteInt(5, 20)

bitEn.WriteBool(false)

bitEn.WriteBool(true)


data := bitEn.GetContent()


bitDe := byteuti.NewBitDecoder(bytes.NewReader(data))

val, _ := bitDe.ReadInt(4) // val is 10

b, _ := bitDe.ReadBool() // b is true

val, _ = bitDe.ReadInt(5) // val is 20

b, _ = bitDe.ReadBool() // b is false

b, _ = bitDe.ReadBoo() // b is true



## Binary and Hex convert

data := []byte{1, 2, 3, 4, 10, 11, 255}

hex := byteutil.BinaryToHex(data) // hex is 010203040A0BFF

data = byteutil.HexToBinary(hex) // data is 1, 2, 3, 4, 10, 11, 255


