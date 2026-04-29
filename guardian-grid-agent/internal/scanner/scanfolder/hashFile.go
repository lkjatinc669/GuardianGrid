package scanfolder

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func hashFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	hasher := sha256.New()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
