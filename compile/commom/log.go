package commom

import (
	"fmt"
	"strconv"
)

type Log interface {
	Log()
}

func LogList[T any](arr []T) {
	fmt.Println("size: " + strconv.Itoa(len(arr)))
	for i := 0; i < len(arr); i++ {
		fmt.Printf("  ")
		fmt.Println(arr[i])
	}
}
