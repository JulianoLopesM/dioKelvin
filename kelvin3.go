package main

import "fmt"

const ebulicaoK float64 = 373

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273

	fmt.Printf("a temperatura de ebulição da água em K =  %g, %T , e a temperatura de ebulição da agua em C =  %g, %T.       ", tempK, tempK, tempC, tempC)

}