package bcrypt

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	pwd := []byte("123456")
	encrypted, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		t.Error(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, pwd)
	assert.NoError(t, err)
	t.Log(string(encrypted))
}
