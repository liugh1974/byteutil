package byteutil

import (
	"bytes"
	"fmt"
	"math"
)

var uintMaxVal = map[uint]uint{
	1: 0xFF,
	2: 0xFFFF,
	3: 0xFFFFFF,
	4: math.MaxUint32,
}

var uint64MaxVal = map[uint]uint64{
	1: 0xFF,
	2: 0xFFFF,
	3: 0xFFFFFF,
	4: 0xFFFFFFFF,
	5: 0xFFFFFFFFFF,
	6: 0xFFFFFFFFFFFF,
	7: 0xFFFFFFFFFFFFFF,
	8: math.MaxUint64,
}

type ByteEncoder struct {
	buf *bytes.Buffer
}

func NewByteEncoder() *ByteEncoder {
	return &ByteEncoder{buf: &bytes.Buffer{}}
}

func (b *ByteEncoder) WriteInt(size, val uint) error {
	if size < 1 || size > 4 {
		return fmt.Errorf("write int size range 1 -- 4")
	}

	if val > uintMaxVal[size] {
		return fmt.Errorf("%d cannot be greater than max value for size %d (%d)", val, size, uintMaxVal[size])
	}

	data := make([]byte, size)
	var i uint
	for ; i < size; i++ {
		data[int(i)] = byte(val >> (((size - i - 1) * 8) & 0xFF))
	}
	b.buf.Write(data)
	return nil
}

func (b *ByteEncoder) WriteInt64(size uint, val uint64) error {
	if size < 1 || size > 8 {
		return fmt.Errorf("write int64 size range 1 -- 8")
	}

	if val > uint64MaxVal[size] {
		return fmt.Errorf("%d cannot be greater than max value for size %d (%d)", val, size, uint64MaxVal[size])
	}

	data := make([]byte, size)
	var i uint
	for ; i < size; i++ {
		data[i] = byte(val >> (((size - i - 1) * 8) & 0xFF))
	}
	b.buf.Write(data)
	return nil
}

func (b *ByteEncoder) WriteBytes(val []byte) error {
	b.buf.Write(val)
	return nil
}

func (b *ByteEncoder) WriteString(val string) error {
	b.buf.Write([]byte(val))
	return nil
}

func (b *ByteEncoder) WriteStringWithSize(size uint, val string) error {
	length := uint(len(val))
	if size < length {
		return fmt.Errorf("string length %d is greater than size: %d", len(val), size)
	}

	if size > length {
		b.buf.Write(make([]byte, int(size-length)))
	}
	b.buf.Write([]byte(val))
	return nil
}

func (b *ByteEncoder) GetContent() []byte {
	return b.buf.Bytes()
}
