package byteutil

import (
	"errors"
	"fmt"
	"strconv"
)

func BinaryToHex(data []byte) string {
	var str string
	for _, b := range data {
		str += fmt.Sprintf("%02X", b)
	}
	return str
}

func HexToBinary(hex string) ([]byte, error) {
	if len(hex)%2 != 0 {
		return nil, errors.New("Error HEX String")
	}

	data := make([]byte, len(hex)/2)
	index := 0
	for i := 0; i < len(hex); i += 2 {
		n, err := strconv.ParseInt(string(hex[i:i+2]), 16, 16)
		if err != nil {
			return nil, err
		}
		data[index] = byte(n)
		index++
	}
	return data, nil
}
