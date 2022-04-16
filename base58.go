package uuid

import (
	"github.com/zergon321/base58"
)

// MarshalBase58 returns the Bitcoin base58
// representation of the UUID.
func (id UUID) MarshalBase58() ([]byte, error) {
	var buffer [22]byte
	err := base58.EncodeToBuffer(id[:], buffer[:])

	if err != nil {
		return nil, err
	}

	return buffer[:], nil
}

// UnmarshalBase58 returns a new UUID from
// the given Bitcoin base58 string.
func UnmarshalBase58(idStr string) (UUID, error) {
	var buffer [34]byte
	var auxBuffer [6]uint32
	start, end, err := base58.DecodeToBufferWithAux(
		idStr, buffer[:], auxBuffer[:])

	if err != nil {
		return UUID{}, err
	}

	dec := buffer[start:end]
	var id UUID

	copy(id[:], dec)

	return id, nil
}
