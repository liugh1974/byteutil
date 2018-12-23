package byteutil

import (
	"fmt"
	"io"
)

type ByteDecoder struct {
	r io.Reader
}

func NewByteDecoder(r io.Reader) *ByteDecoder {
	return &ByteDecoder{r: r}
}

func (b *ByteDecoder) ReadInt(size uint) (uint, error) {
	if size < 1 || size > 4 {
		return 0, fmt.Errorf("read size %d is not in int size range 1 -- 4", size)
	}

	buf := make([]byte, size)
	n, err := b.r.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != int(size) {
		return 0, fmt.Errorf("Read exception, should be read %d, actually read %d", size, n)
	}

	var val uint
	for _, d := range buf {
		val = (val << 8) | uint(d)
	}
	return val, nil
}

func (b *ByteDecoder) ReadInt64(size uint) (uint64, error) {
	if size < 1 || size > 8 {
		return 0, fmt.Errorf("read size %d is not in int64 size range 1 -- 8", size)
	}

	buf := make([]byte, size)
	n, err := b.r.Read(buf)
	if err != nil {
		return 0, err
	}
	if n != int(size) {
		return 0, fmt.Errorf("Read exception, should be read %d, actually read %d", size, n)
	}

	var val uint64
	for _, d := range buf {
		val = (val << 8) | uint64(d)
	}
	return val, nil
}

func (b *ByteDecoder) ReadBytes(size uint) ([]byte, error) {
	buf := make([]byte, size)
	n, err := b.r.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != int(size) {
		return nil, fmt.Errorf("Read exception, should be read %d, actually read %d", size, n)
	}
	return buf, nil
}

func (b *ByteDecoder) ReadString() (string, error) {
	data, err := b.ReadRemains()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (b *ByteDecoder) ReadStringWithSize(size uint) (string, error) {
	data, err := b.ReadBytes(size)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func (b *ByteDecoder) HasRemain() bool {
	data, _ := b.ReadRemains()
	return len(data) > 0
}

func (b *ByteDecoder) ReadRemains() ([]byte, error) {
	data := []byte{}
	for {
		buf := make([]byte, 1024)
		n, err := b.r.Read(buf)
		if err != nil {
			return nil, err
		}

		if n < 1024 {
			data = append(data, buf[:n]...)
			break
		}
		data = append(data, buf...)
	}
	return data, nil
}
