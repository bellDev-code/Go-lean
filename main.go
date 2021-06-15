package main

import "fmt"


func main() {
	array := map[string]string{"name":"lee", "age":"27"}
	for key, value := range array {
		fmt.Println(key, value)
	}
	
}