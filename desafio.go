package main

import "fmt"

const ebulicaoK = 373.2

func main(){

	k := ebulicaoK
	c := (k - 273) * 5 / 9

	fmt.Println("A temperatura de ebuliçao da agua em ºK =", k)
	fmt.Println("A tempretaura de ebulição da água em ºC =", c)
}