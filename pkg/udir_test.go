package pkg

import (
	"fmt"
	"testing"
)

func Test_ListFileInDir(t *testing.T) {
	for _, f := range ListFileInDir("/users/clpsz/Desktop/", false) {
		fmt.Println(f)
	}
}

func Test_ListSubDirInDir(t *testing.T) {
	subs := ListSubDirInDir("/users/clpsz/Desktop/", true)
	for _, d := range subs {
		fmt.Println(d)
	}
}
