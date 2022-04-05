package main

import (
	api "./gin"
	"fmt"
)

func main() {
	fmt.Println("Run Server")
	api.RunGinExample()
	//api.RunQueryParam()
}
