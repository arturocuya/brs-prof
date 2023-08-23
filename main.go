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
	ActualSampleRatio     float32
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

func main() {
	file, err := os.Open("./input.bsprof")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := Header{}

	reader := bufio.NewReader(file)

	// Discard magic bytes
	_, err = reader.Read(make([]byte, 8))
	if err != nil {
		panic(err)
	}

	if err = ReadUInt32(reader, &header.Major); err != nil {
		panic(err)
	}

	if err = ReadUInt32(reader, &header.Minor); err != nil {
		panic(err)
	}

	if err = ReadUInt32(reader, &header.Patch); err != nil {
		panic(err)
	}

	if err = ReadUInt32(reader, &header.Size); err != nil {
		panic(err)
	}

	if err = ReadFloat32(reader, &header.RequestedSampleRatio); err != nil {
		panic(err)
	}

	if err = ReadFloat32(reader, &header.ActualSampleRatio); err != nil {
		panic(err)
	}

	if err = ReadBool(reader, &header.IncludesLineData); err != nil {
		panic(err)
	}

	if err = ReadBool(reader, &header.IncludesMemOps); err != nil {
		panic(err)
	}

	var timestampUInt uint64
	if err = ReadUInt64(reader, &timestampUInt); err != nil {
		panic(err)
	}
	header.StartTime = time.Unix(int64(timestampUInt), 0) 

	if err = ReadUtf8z(reader, &header.ChannelName); err != nil {
		panic(err)
	}

	if err = ReadUtf8z(reader, &header.SupplementalInfo); err != nil {
		panic(err)
	}

	if err = ReadUtf8z(reader, &header.ChannelVersion); err != nil {
		panic(err)
	}

	if err = ReadUtf8z(reader, &header.DeviceVendorName); err != nil {
		panic(err)
	}

	if err = ReadUtf8z(reader, &header.DeviceModelNumber); err != nil {
		panic(err)
	}

	if err = ReadUtf8z(reader, &header.DeviceFirmwareVersion); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", header)
}
