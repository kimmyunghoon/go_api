package main

import (
	"./driver"
	"fmt"
)

func main() {
	driver.MongoDB()
	fmt.Println("Run Server")
	//api.RunGinExample()
	//api.RunQueryParam()
}
