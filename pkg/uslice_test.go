package pkg

import (
	"fmt"
	"testing"
)

func Test_ChunkIntSlice(t *testing.T) {
	fmt.Println(ChunkIntSlice([]int{1,2,3,4,5,6,7,8,9,10}, 3))
	fmt.Println(ChunkIntSlice([]int{1,2,3,4,5,6,7,8,9,10}, 4))
	fmt.Println(ChunkIntSlice([]int{1,2,3,4,5,6,7,8,9,10}, 5))
	fmt.Println(ChunkIntSlice([]int{1,2,3,4,5,6,7,8,9,10}, 10))
	fmt.Println(ChunkIntSlice([]int{1,2,3,4,5,6,7,8,9,10}, 11))
}
