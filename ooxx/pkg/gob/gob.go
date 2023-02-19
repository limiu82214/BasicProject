package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/pkg/errors"
)

func StoreStructToByte(a any) ([]byte, error) {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(a)

	return data.Bytes(), err
}

func GetStrutFromByte(b []byte, a any) error {
	enc := gob.NewDecoder(bytes.NewReader(b))
	err := enc.Decode(a)

	return errors.Wrap(err, ":GetStrutFromByte")
}
