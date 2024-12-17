package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func exercises1_1(s []string) {
	fmt.Println(s[0])
}

func exercises1_2(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}

func exercises1_3Joint(s []string) {
	fmt.Println(strings.Join(s, " "))
}

func main() {

	s := os.Args //example, $ go run exercises1.1_1.3.go -f -d

	exercises1_1(s)

	exercises1_2(s)

	start := time.Now()
	exercises1_2(s)
	end1 := time.Since(start).Microseconds()

	start = time.Now()
	exercises1_3Joint(s)
	end2 := time.Since(start).Microseconds()

	fmt.Printf("Usual: %v, help with strings.Join: %v", end1, end2) //approximately Usual: 528, with strings.Join: 0

}
