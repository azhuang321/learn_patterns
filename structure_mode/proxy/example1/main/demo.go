package main

import (
	"fmt"
	"go-patterns/structure_mode/proxy"
)

func main() {
	image := &proxy.ProxyImage{
		FileName: "test.png",
	}
	//image 将从磁盘加载
	image.Display()
	fmt.Println()
	//image 不会从磁盘加载
	image.Display()
}
