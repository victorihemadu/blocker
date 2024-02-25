package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := generatePrivateKey()

	assert.Equal(t, len(privKey.Bytes()), privKeyLen)
}