package uuid_test

import (
	"testing"

	b58 "github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/zergon321/uuid"
)

func TestBase58Encode(t *testing.T) {
	id := uuid.New()

	b58Str, err := id.MarshalBase58()
	assert.Nil(t, err)

	b58AltStr := b58.Encode(id[:])
	assert.Equal(t, b58AltStr, b58Str)
}

func TestBase58Decode(t *testing.T) {
	str := "YaUV8Ysvm9EjpT8mmbTbwU"

	id, err := uuid.UnmarshalBase58(str)
	assert.Nil(t, err)

	idAltData, err := b58.Decode(str)
	assert.Nil(t, err)

	var idAlt uuid.UUID
	copy(idAlt[:], idAltData)
	assert.Equal(t, id, idAlt)
}
