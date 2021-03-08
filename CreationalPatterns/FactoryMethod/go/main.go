package main

import "fmt"

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

func main()  {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	_, res := getGun("test")
	fmt.Println(res.Error())
	printDetails(ak47)
	printDetails(musket)
}
