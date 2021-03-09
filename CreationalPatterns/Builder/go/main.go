package main

import "fmt"

func main()  {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Type: %d\n", normalHouse.floor)

	director.setBuild(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("igloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("igloo House Num Type: %d\n", iglooHouse.floor)
}