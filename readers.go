package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
)

func ReadUInt64(reader *bufio.Reader, target *uint64) error {
	var buf [8]byte

	_, err := io.ReadFull(reader, buf[:])
	if err != nil {
		return err
	}

	*target = binary.LittleEndian.Uint64(buf[:])
	return nil
}

func ReadUInt32(reader *bufio.Reader, target *uint32) error {
	var shift uint

	for {
		// Read a byte
		b, err := reader.ReadByte()
		if err != nil {
			return err
		}

		// Extract the 7 lower bits
		value := uint32(b & 0x7F)
		*target |= value << shift

		// Check if the high bit is set
		if (b & 0x80) == 0 {
			break
		}

		shift += 7
	}

	return nil
}

func ReadFloat32(reader *bufio.Reader, target *float32) error {
	var bits uint32
	err := ReadUInt32(reader, &bits)
	if err != nil {
		return err
	}
	*target = math.Float32frombits(bits)
	return nil
}

func ReadBool(reader *bufio.Reader, target *bool) error {
	b, err := reader.ReadByte()
	if err != nil {
		return err
	}
	*target = b != 0 
	return nil
}

func ReadUtf8z(reader *bufio.Reader, target *string) error {
	var stringBuffer bytes.Buffer

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return errors.New("reached EOF before finding null terminator")
			}
			return err
		}

		if char == 0 {
			break
		}

		stringBuffer.WriteRune(char)
	}

	*target = stringBuffer.String()
	return nil
}

