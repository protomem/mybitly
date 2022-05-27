package crypt

import (
	"crypto/sha256"
	"math/big"

	"github.com/speps/go-hashids/v2"
)

const (
	_salt = "Eiweoinef32392fmfesslk"
)

func Sha256Of(input string) []byte {

	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum([]byte(_salt))

}

func GenerateShortToken(data string) (string, error) {

	// TODO Refactor

	hd := hashids.NewData()
	hd.Salt = _salt
	hd.MinLength = 4

	hashData := Sha256Of(data)
	numb := new(big.Int).SetBytes(hashData).Int64()

	h, _ := hashids.NewWithData(hd)
	shortToken, err := h.EncodeInt64([]int64{numb})

	return shortToken, err

}
