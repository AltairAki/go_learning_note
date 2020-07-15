package loop_test

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	/*while(n<5)*/
	for n < 5 {
		n++
		fmt.Print(" ", n)
	}
	/*while(true)*/
	// for {
	// 	fmt.Println(n)
	// }
}
