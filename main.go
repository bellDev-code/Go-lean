package main

import "fmt"

func reapeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	reapeatMe("jong", "ho", "young", "do")
}