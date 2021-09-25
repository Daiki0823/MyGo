package main

import (
	"fmt"
)

func main() {
	//for文版
	// for i := 0; i <= 100; i = i + 1 {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%dは偶数\n", i)

	// 	} else {
	// 		fmt.Printf("%dは奇数\n", i)
	// 	}
	// }

	//switch文版
	for i := 0; i <= 100; i++ {
		switch i % 2 {
		case 0:
			fmt.Printf("%dは偶数\n", i)
		case 1:
			fmt.Printf("%dは奇数\n", i)
		}

	}

}
