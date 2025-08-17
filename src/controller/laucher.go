package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

type LaucherController struct{}

func (lc *LaucherController) HashFile(path string) (int64, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, "", err
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return 0, "", err
	}
	return n, hex.EncodeToString(h.Sum(nil)), nil
}
