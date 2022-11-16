package main 

import "fmt"

func main() {
	suma := Suma(2,2)
	fmt.Println("El resultado de la suma es: ", suma)
	
	resta := Resta(2,2)
	fmt.Println("El resultado de la resta es: ", resta)
	
	multiplicacion := Multiplicacion(2,2)
	fmt.Println("El resultado de la multiplicacion es: ", multiplicacion)
	
	division := Division(2,2)
	fmt.Println("El resultado de la division es: ", division)
	
}

func Suma(numero1, numero2 int) (resultado int){
	resultado = numero1 + numero2 
	return
}

func Resta(numero1, numero2 int) (resultado int){
	resultado = numero1 - numero2 
	return
}

func Multiplicacion(numero1, numero2 int) (resultado int){
	resultado = numero1 * numero2 
	return
}

func Division(numero1, numero2 int) (resultado int){
	resultado = numero1 / numero2 
	return
}
