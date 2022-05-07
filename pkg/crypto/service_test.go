package crypto

import (
	"testing"

	"github.com/mchmarny/twcli/pkg/config"
	"github.com/stretchr/testify/assert"
)

const (
	badStr    = "not-a-valid-base64-string"
	clearText = "some-secret-value"
)

func TestCrypto(t *testing.T) {
	c, err := config.ReadFromFile("../../configs/unit.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, c.Crypto)
	assert.NotEmpty(t, c.Crypto.EncryptionKey)

	t.Run("new without key", func(t *testing.T) {
		_, err := New("")
		assert.Error(t, err)
	})
	t.Run("new with an invalid key", func(t *testing.T) {
		_, err := New("test")
		assert.Error(t, err)
	})
	t.Run("new with valid key", func(t *testing.T) {
		s, err := New(c.Crypto.EncryptionKey)
		assert.NoError(t, err)
		assert.NotNil(t, s)
	})

	t.Run("full crypt test", func(t *testing.T) {
		s1, err := New(c.Crypto.EncryptionKey)
		assert.NoError(t, err)
		assert.NotNil(t, s1)

		e1, err := s1.Encrypt(clearText)
		assert.NoError(t, err)
		assert.NotEmpty(t, e1)

		d1, err := s1.Decrypt(e1)
		assert.NoError(t, err)
		assert.NotEmpty(t, e1)
		assert.Equal(t, clearText, d1)

		e2, err := s1.Encrypt(clearText)
		assert.NoError(t, err)
		assert.NotEmpty(t, e2)
		assert.Equal(t, e1, e2)

		d2, err := s1.Decrypt(e2)
		assert.NoError(t, err)
		assert.NotEmpty(t, e2)
		assert.Equal(t, clearText, d2)
		assert.Equal(t, d1, d2)

		_, err = s1.Decrypt(badStr)
		assert.Error(t, err)
	})
	t.Run("encoding/decodding", func(t *testing.T) {
		v := encode([]byte(clearText))
		assert.NotEmpty(t, v)
		b, err := decode(v)
		assert.NoError(t, err)
		assert.NotEmpty(t, b)
		assert.Equal(t, clearText, string(b))
		_, err = decode(badStr)
		assert.Error(t, err)
	})
}
