package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Hash     string
	Password string
	Expected bool
}

func TestGenerateHash(t *testing.T) {
	password := "Github_:)"
	hashedString, err := GenerateHash(password)
	fmt.Println(hashedString)
	assert.Nil(t, err)
	assert.NotEqual(t, "", hashedString)
}

func TestCheckHash(t *testing.T) {
	var tests []TestStruct = make([]TestStruct, 2)
	tests[0] = TestStruct{
		Password: "Github_:)",
		Hash:     "$2a$10$6iBQpnvfYWCpH1JUyJsLZeiPBihkEm6.TOPw8NNXlc49J1082cYmq",
		Expected: true,
	}
	tests[1] = TestStruct{
		Password: "wrong-password",
		Hash:     "$2a$10$6iBQpnvfYWCpH1JUyJsLZeiPBihkEm6.TOPw8NNXlc49J1082cYmq",
		Expected: false,
	}

	for _, v := range tests {
		result := CheckHash(v.Hash, v.Password)
		assert.Equal(t, v.Expected, result)
	}
}
