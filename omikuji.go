package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceSeed := rand.Intn(6)
	//switch文版

	switch diceSeed + 1 {
	case 1:
		fmt.Println("君の運勢は凶です、頑張ろう、、、")
	case 2, 3:
		fmt.Println("運勢は吉だよ！何とも言えないね！")
	case 4, 5:
		fmt.Println("運勢は中吉だよ！いい感じ！")
	default:
		fmt.Println("さすがだよ！大吉だ！")
	}

}
