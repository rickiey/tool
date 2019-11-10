package tool

import (
	"bytes"
	"encoding/binary"
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateUUIDWithOutLine() string {
	u := uuid.New().String()

	return strings.ReplaceAll(u, "-", "")
}

//
func GenerateUUID2Number() uint64 {
	var res uint64
	newuuid := uuid.New()
	bf, _ := newuuid.MarshalBinary()
	buf := bytes.NewReader(bf)
	_ = binary.Read(buf, binary.LittleEndian, &res)
	return res
}
