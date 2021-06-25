package pkg

import (
	"testing"
)

func Test_GetRandomString(t *testing.T) {
	if len(GetRandomString(10)) != 10 {
		t.Fatalf("string length error\n")
	}
}
