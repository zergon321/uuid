package uuid_test

import (
	"testing"

	"github.com/zergon321/uuid"
)

func BenchmarkBase58Encode(b *testing.B) {
	id := uuid.New()

	for i := 0; i < b.N; i++ {
		id.MarshalBase58()
	}
}

func BenchmarkStandardEncode(b *testing.B) {
	id := uuid.New()

	for i := 0; i < b.N; i++ {
		id.MarshalText()
	}
}
