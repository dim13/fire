package main

import (
	"crypto/rand"
	"encoding/binary"
)

func cryptoSeed() int64 {
	var x int64
	binary.Read(rand.Reader, binary.LittleEndian, &x)
	return x
}
