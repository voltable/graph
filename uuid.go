// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graph

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const uuidSize = 16

func generateRandomBytes() ([uuidSize]byte, error) {
	buf := make([]byte, uuidSize)
	var arr [uuidSize]byte
	if _, err := rand.Read(buf); err != nil {
		return arr, fmt.Errorf("failed to read random bytes: %v", err)
	}
	copy(arr[:], buf)
	return arr, nil
}

func formatUUID(buf [uuidSize]byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16])
}

func parseUUID(uuid string) ([]byte, error) {
	if len(uuid) != 36 {
		return nil, fmt.Errorf("uuid string is wrong length")
	}

	hyph := []byte("-")

	if uuid[8] != hyph[0] ||
		uuid[13] != hyph[0] ||
		uuid[18] != hyph[0] ||
		uuid[23] != hyph[0] {
		return nil, fmt.Errorf("uuid is improperly formatted")
	}

	hexStr := uuid[0:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:36]

	ret, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	if len(ret) != uuidSize {
		return nil, fmt.Errorf("decoded hex is the wrong length")
	}

	return ret, nil
}
