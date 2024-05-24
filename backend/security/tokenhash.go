package security

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/twinj/uuid"
)

func TokenHash(text string) string {
	hashes := md5.New()
	hashes.Write([]byte(text))
	theHash := hex.EncodeToString(hashes.Sum(nil))

	//also use uuid
	u := uuid.NewV4()
	theToken := theHash + u.String()

	return theToken
}
