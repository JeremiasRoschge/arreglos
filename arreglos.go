package main

import (
	"fmt"
)

func main()  {
	
	arreglo := [3]int{4,2}

	fmt.Println("Espacios del arreglo:", len(arreglo))

	var matriz [3][2]int
	matriz[0][0] = 10
	matriz[0][1] = 19 

	fmt.Println(matriz)
	fmt.Println(matriz[0][1])
	fmt.Println(len(matriz[0]))
}