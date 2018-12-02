package byteutil

import "fmt"

type BitDecoder struct {
	stream  []byte
	offset  uint
	bitLeft uint
	total   uint
}

func NewBitDecoder(data []byte) *BitDecoder {
	return &BitDecoder{data, 0, 8, uint(len(data) * 8)}
}

func (b *BitDecoder) ReadInt(size uint) (uint, error) {
	if b.total < size {
		return 0, fmt.Errorf("end of stream for reading %d bits", size)
	}
	b.total -= size
	var val uint
	for size > 0 {
		count := Min(size, b.bitLeft)
		bits := b.stream[b.offset] >> (b.bitLeft - count) & ((1 << count) - 1)
		val = (val << count) | uint(bits)
		size -= count
		b.bitLeft -= count
		if b.bitLeft <= 0 {
			b.bitLeft = 8
			b.offset++
		}
	}
	return val, nil
}

func (b *BitDecoder) ReadBool() (bool, error) {
	val, err := b.ReadInt(1)
	if err != nil {
		return false, err
	}
	return val == 1, nil
}
