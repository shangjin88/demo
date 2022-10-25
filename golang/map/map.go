package main

import "fmt"

func main() {
	mapHaiCoder := map[string]string{
		"Server":     "Golang",
		"JavaScript": "Vue",
		"Db":         "Redis",
	}
	value, isOk := mapHaiCoder["Server"]
	fmt.Println("Value =", value, "IsOk =", isOk)
}
