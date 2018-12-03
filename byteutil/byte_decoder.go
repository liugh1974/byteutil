package byteutil

import (
	"fmt"
)

type ByteDecoder struct {
	stream []byte
	offset uint
	total  uint
}

func NewByteDecoder(data []byte) *ByteDecoder {
	return &ByteDecoder{data, 0, uint(len(data))}
}

func (b *ByteDecoder) ReadInt(size uint) (uint, error) {
	if size < 1 || size > 4 {
		return 0, fmt.Errorf("read size %d is not in int size range 1 -- 4", size)
	}
	if b.offset+size > b.total {
		return 0, fmt.Errorf("read int end of stream for size: %d", size)
	}

	var val uint
	var i uint = 0
	for ; i < size; i++ {
		val = val << 8
		val += uint(b.stream[b.offset])
		b.offset++
	}
	return val, nil
}

func (b *ByteDecoder) ReadInt64(size uint) (uint64, error) {
	if size < 1 || size > 8 {
		return 0, fmt.Errorf("read size %d is not in int64 size range 1 -- 8", size)
	}
	if b.offset+size > b.total {
		return 0, fmt.Errorf("read int64 end of stream for size: %d", size)
	}

	var val uint64
	var i uint = 0
	for ; i < size; i++ {
		val = val << 8
		val += uint64(b.stream[b.offset])
		b.offset++
	}
	return val, nil
}

func (b *ByteDecoder) ReadBytes(size uint) ([]byte, error) {
	if b.offset+size > b.total {
		return nil, fmt.Errorf("read bytes end of stream for size: %d", size)
	}
	start := b.offset
	b.offset += size
	return b.stream[start:b.offset], nil
}

func (b *ByteDecoder) ReadString() (string, error) {
	re, _ := b.ReadRemains()
	return string(re), nil
}

func (b *ByteDecoder) ReadStringWithSize(size uint) (string, error) {
	re, err := b.ReadBytes(size)
	if err != nil {
		return "", err
	}
	return string(re), nil

}

func (b *ByteDecoder) HasRemain() bool {
	return b.offset < b.total
}

func (b *ByteDecoder) ReadRemains() ([]byte, error) {
	return b.stream[b.offset:], nil
}
