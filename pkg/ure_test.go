package pkg

import (
	"fmt"
	"testing"
)

func Test_GetReGroupMany(t *testing.T) {
	fmt.Println(GetReGroupMany(
		`^A (wei\w+) (fell).*(l[^ ]+).*(C\w+).$`,
		"A weight fell like a stone upon Clym.",
		[]int{1,3,4}),
	)
}
