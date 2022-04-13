package main

import (
	"crypto/rand"
	"encoding/binary"
)

func cryptoSeed() (x int64) {
	binary.Read(rand.Reader, binary.LittleEndian, &x)
	return
}
