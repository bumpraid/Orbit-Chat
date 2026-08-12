package main

import "fmt"

var status = false
func init() {
	fmt.Println("Started...")
}
func main() {
	fmt.Println("Ok")
	status = true
	if status {
		fmt.Println("Status: Online")
	}
}

