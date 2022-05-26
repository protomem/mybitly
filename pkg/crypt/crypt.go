package crypt

import (
	"github.com/speps/go-hashids/v2"
)

func GenerateShortToken(data string) (string, error) {

	// TODO Refactor

	hd := hashids.NewData()
	hd.Salt = data
	hd.MinLength = 6
	h, _ := hashids.NewWithData(hd)
	shortToken, err := h.Encode([]int{45, 434, 1313, 99})

	return shortToken, err

}
