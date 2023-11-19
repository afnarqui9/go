package main

import (
	"fmt"
)

func main(){
	// const pi float64 = 3.14
	// const pi2 = 3.14
	// fmt.Println("pi: ", pi)
	// fmt.Println("pi2: ", pi2)

	// base := 12
	// var altura int = 14
	// var area  int

	// fmt.Println("base: ", base)
	// fmt.Println("altura: ", altura)
	// fmt.Println("area: ", area)

	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a,b,c,d)

//Numeros enteros
//int = Depende del OS (32 o 64 bits)
//int8 = 8bits = -128 a 127
//int16 = 16bits = -2^15 a 2^15-1
//int32 = 32bits = -2^31 a 2^31-1
//int64 = 64bits = -2^63 a 2^63-1

//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
//uint = Depende del OS (32 o 64 bits)
//uint8 = 8bits = 0 a 127
//uint16 = 16bits = 0 a 2^15-1
//uint32 = 32bits = 0 a 2^31-1
//uint64 = 64bits = 0 a 2^63-1

//numeros decimales
// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

//textos y booleanos
//string = ""
//bool = true or false

//numeros complejos
//Complex64 = Real e Imaginario float32
//Complex128 = Real e Imaginario float64
//Ejemplo : c:=10 + 8i	
nombre := "vienen"
valores := 100
fmt.Printf("%s muchos afnarqui %d aja \n", nombre, valores)
message := fmt.Sprintf("%s muchos %d aja \n", nombre, valores)
fmt.Println(message)
fmt.Printf("type: %T", valores)

}