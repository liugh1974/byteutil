package byteutil

import "fmt"

type BitEncoder struct {
	stream  []byte
	offset  uint
	bitLeft uint
}

func NewBitEncoder() BitEncoder {
	return BitEncoder{make([]byte, 4), 0, 8}
}

func (b *BitEncoder) GetContent() []byte {
	end := b.offset
	if b.bitLeft != 8 {
		end++
	}
	return b.stream[:end]
}

func (b *BitEncoder) WriteInt(size, val uint) error {
	if size < 1 {
		return fmt.Errorf("wrinte int size cannot be less than 1")
	}

	if err := checkValue(size, val); err != nil {
		return err
	}

	for size > 0 {
		count := Min(size, b.bitLeft)
		bits := (val >> (size - count)) & ((1 << count) - 1)
		b.stream[b.offset] = b.stream[b.offset] | byte(bits<<(b.bitLeft-count))
		size -= count
		b.bitLeft -= count
		if b.bitLeft <= 0 {
			b.bitLeft = 8
			b.offset++
			b.resize()
		}
	}

	return nil
}

func (b *BitEncoder) WriteBool(v bool) error {
	val := 0
	if v {
		val = 1
	}
	return b.WriteInt(1, uint(val))
}

func (b *BitEncoder) resize() {
	if b.offset >= uint(len(b.stream)) {
		buf := make([]byte, len(b.stream)*2)
		b.stream = append(b.stream, buf...)
	}
}

func checkValue(size, val uint) error {
	if val > ((1 << size) - 1) {
		return fmt.Errorf("%d is  too big with %b bits", val, size)
	}
	return nil
}

func Min(n, m uint) uint {
	if n > m {
		return m
	} else {
		return n
	}
}
