
// +build integration
 
package main
 
import "testing"
import "fmt"
            
func TestIntegrationCalculadora(t *testing.T) {
            suma := Suma(8, 8)
            resta := Resta(8, 8)     
            multiplicacion := Multiplicacion(8, 8)
            division := Division(8, 8)
            promedio := Promedio(8, 8, 8, 8, 8)
 
            promediototal := Promedio(suma, resta, multiplicacion, division, promedio)
 
            if promediototal != 17 {
                        t.Error("Se esperaba 17 y se obtuvo", promediototal)
            }
            
            fmt.Println("Prueba de integracion del promedio total fue exitosa")
}
