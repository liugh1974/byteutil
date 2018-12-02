package byteutil

import (
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
	stream []byte
	offset uint
}

func NewByteEncoder() *ByteEncoder {
	return &ByteEncoder{make([]byte, 128), 0}
}

func (b *ByteEncoder) WriteInt(size, val uint) error {
	if size < 1 || size > 4 {
		return fmt.Errorf("write int size range 1 -- 4")
	}

	if val > uintMaxVal[size] {
		return fmt.Errorf("%d cannot be greater than max value for size %d (%d)", val, size, uintMaxVal[size])
	}

	b.resize(size)
	var i uint = 0

	for ; i < size; i++ {
		b.stream[b.offset] = byte(val >> (((size - i - 1) * 8) & 0xFF))
		b.offset++
	}
	return nil
}

func (b *ByteEncoder) WriteInt64(size uint, val uint64) error {
	if size < 1 || size > 8 {
		return fmt.Errorf("write int64 size range 1 -- 8")
	}

	if val > uint64MaxVal[size] {
		return fmt.Errorf("%d cannot be greater than max value for size %d (%d)", val, size, uint64MaxVal[size])
	}

	b.resize(size)

	var i uint = 0
	for ; i < size; i++ {
		b.stream[b.offset] = byte(val >> (((size - i - 1) * 8) & 0xFF))
		b.offset++
	}
	return nil
}

func (b *ByteEncoder) WriteBytes(val []byte) error {
	b.stream = append(b.stream[:b.offset], val...)
	b.offset += uint(len(val))
	return nil
}

func (b *ByteEncoder) WriteString(val string) error {
	return b.WriteStringWithSize(uint(len(val)), val)
}

func (b *ByteEncoder) WriteStringWithSize(size uint, val string) error {
	if size < uint(len(val)) {
		return fmt.Errorf("string length %d is greater than size: %d", len(val), size)
	}

	spaceSize := size - uint(len(val))
	if spaceSize > 0 {
		spaces := make([]byte, spaceSize)
		b.stream = append(b.stream[:b.offset], spaces...)
		b.offset += uint(spaceSize)
	}

	b.stream = append(b.stream[:b.offset], []byte(val)...)
	b.offset += uint(len(val))
	return nil
}

func (b *ByteEncoder) resize(size uint) {
	size = b.offset + size
	if size < uint(len(b.stream)*2) {
		size = uint(len(b.stream) * 2)
	} else {
		size = size + 1
	}
	buf := make([]byte, size)
	b.stream = append(b.stream, buf...)
}

func (b *ByteEncoder) GetContent() []byte {
	return b.stream[:b.offset]
}
