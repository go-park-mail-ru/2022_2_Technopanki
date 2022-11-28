package codeGenerator

import (
	"crypto/rand"
	"math/big"
)

func GenerateCode() (string, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	r.Add(r, big.NewInt(100000))
	return r.String(), nil
}
