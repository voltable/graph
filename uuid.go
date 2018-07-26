package graph

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// UUID
type UUID [16]byte

func GenerateRandomUUID() (UUID, error) {
	buf := make([]byte, 16)
	var arr [16]byte
	if _, err := rand.Read(buf); err != nil {
		return arr, fmt.Errorf("failed to read random bytes: %v", err)
	}
	copy(arr[:], buf)
	return arr, nil
}

func sliceToVertexID(buf []byte) UUID {
	var arr [16]byte
	copy(arr[:], buf)
	return arr
}

func formatUUID(buf [16]byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}

func parseUUID(uuid string) ([16]byte, error) {
	var arr [16]byte
	if len(uuid) != 36 {
		return arr, fmt.Errorf("uuid string is wrong length")
	}

	hyph := []byte("-")

	if uuid[8] != hyph[0] ||
		uuid[13] != hyph[0] ||
		uuid[18] != hyph[0] ||
		uuid[23] != hyph[0] {
		return arr, fmt.Errorf("uuid is improperly formatted")
	}

	hexStr := uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:36]

	ret, err := hex.DecodeString(hexStr)
	if err != nil {
		return arr, err
	}
	if len(ret) != 16 {
		return arr, fmt.Errorf("decoded hex is the wrong length")
	}

	copy(arr[:], ret)
	return arr, nil
}
