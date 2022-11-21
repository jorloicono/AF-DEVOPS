
package main
            
import "fmt"
 
func main() {
            suma := Suma(2, 2)
            fmt.Println("El resultado de la suma es: ", suma)
            
            resta := Resta(2, 2)
            fmt.Println("El resultado de la resta es: ", resta)
            
            multiplicacion := Multiplicacion(2, 2)
            fmt.Println("El resultado de la multiplicacion es: ", multiplicacion)
 
            division := Division(2, 2)
            fmt.Println("El resultado de la division es: ", division)
 
            promedio  := Promedio(2, 2, 2, 2, 2)
            fmt.Println("El resultado del promedio es: ", promedio)
}
 
func Suma(numero1, numero2 int) (resultado int) {
            resultado = numero1 + numero2
            return
}
 
func Resta(numero1, numero2 int) (resultado int) {
            resultado = numero1 - numero2
            return
}
 
func Multiplicacion(numero1, numero2 int) (resultado int) {
            resultado = numero1 * numero2
            return
}
 
func Division(numero1, numero2 int) (resultado int) {
            resultado = numero1 / numero2
            return
}
 
func Promedio(numero1, numero2, numero3, numero4, numero5 int) (resultado int) {
            resultado = (numero1 + numero2 + numero3 + numero4 + numero5) / 5
            return
}

    
