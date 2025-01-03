// // // // // package main

// // // // // import (
// // // // // 	"fmt"
// // // // // )

// // // // // func add(x int, y int) int {
// // // // // 	return x + y
// // // // // }

// // // // //	func main() {
// // // // //		fmt.Println(add(42, 13))
// // // // //	}
// // // // //
// // // // // ==================
// // // // package main

// // // // import "fmt"

// // // // func swap(x, y string) (string, string) {
// // // // 	return y, x
// // // // }

// // // // func main() {
// // // // 	a, b := swap("hello", "world")
// // // // 	fmt.Println(a, b)
// // // // }
// // // // ===============Variable ================

// // // package main

// // // import "fmt"

// // // var c, python, java bool

// // // func main() {
// // // 	// i := 0
// // // 	var i string = "10"
// // // 	fmt.Println(i, c, python, java)
// // // }
// // // =========Datatype====================

// // package main

// // import (
// // 	"fmt"
// // 	"math/cmplx"
// // )

// // var (
// // 	ToBe   bool       = false
// // 	MaxInt uint64     = 1<<64 - 1
// // 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// // )

// // func main() {
// // 	fmt.Printf("Type:%T value:%v\n", ToBe, ToBe)
// // 	fmt.Printf("Type:%T value:%v\n", MaxInt, MaxInt)
// // }
// // ============for loop========
// // package main
// // import "fmt"

// // func main(){
// // 	sum:=0
// // 	for (i:=0;i<10;i++){
// // 		sum=sum+i;

// //		}
// //		fmt.printLn
// //	}
// package main

// import "fmt"

//	func main() {
//		// 	sum := 0
//		// 	for i := 0; i < 10; i++ {
//		// 		sum = sum + i
//		// 	}
//		// 	fmt.Println(sum)
//		i := 42
//		j := 2701
//		p := &i
//		fmt.Println(p)
//		*p = 21
//		fmt.Println(i)
//		p = &j
//		*p = *p / 37
//		fmt.Println(j)
//	}
package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	x := 4
	y := 8
	z := x + y
	fmt.Println(z)
}
