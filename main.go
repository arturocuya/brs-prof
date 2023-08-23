package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Header struct {
	Major                 uint32
	Minor                 uint32
	Patch                 uint32
	Size                  uint32
	RequestedSampleRatio  float32
	ActualSampleRation    float32
	IncludesLineData      bool
	IncludesMemOps        bool
	StartTime             time.Time
	ChannelName           string
	SupplementalInfo      string
	ChannelVersion        string
	DeviceVendorName      string
	DeviceModelNumber     string
	DeviceFirmwareVersion string
}

func readLEB128(reader *bufio.Reader) (uint32, error) {
	var result uint32
	var shift uint

	for {
		// Read a byte
		b, err := reader.ReadByte()
		if err != nil {
			return 0, err
		}

		// Extract the 7 lower bits
		value := uint32(b & 0x7F)
		result |= value << shift

		// Check if the high bit is set
		if (b & 0x80) == 0 {
			break
		}

		shift += 7
	}

	return result, nil
}

func main() {
	file, err := os.Open("./input.bsprof")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := Header{}

	reader := bufio.NewReader(file)

	data := make([]byte, 8)

	bytesRead, err := reader.Read(data)
	if err != nil {
		panic(err)
	}

	text := string(data[:bytesRead])

	major, _ := readLEB128(reader)
	header.Major = major
	minor, _ := readLEB128(reader)
	header.Minor = minor
	build, _ := readLEB128(reader)
	header.Patch = build

	fmt.Println(text)
	fmt.Printf("%+v", header)
}
