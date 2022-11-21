// +build integration
 
package main
 
import "testing"
import "fmt"
import "github.com/joho/godotenv"
import "os"
import "strconv"
            
func TestIntegrationCalculadora(t *testing.T) {
            err := godotenv.Load()
            if err != nil {
            t.Fatalf("could not load .env file: %v", err)
            }
 
            valor_usar, err := strconv.Atoi(os.Getenv("valor_usar"))
            promedio_total, err := strconv.Atoi(os.Getenv("promedio_total"))
 
            suma := Suma(valor_usar, valor_usar)
            resta := Resta(valor_usar, valor_usar)            
            multiplicacion := Multiplicacion(valor_usar, valor_usar)
            division := Division(valor_usar, valor_usar)
            promedio := Promedio(valor_usar, valor_usar, valor_usar, valor_usar, valor_usar)
 
            promediototal := Promedio(suma, resta, multiplicacion, division, promedio)
 
            if promediototal != promedio_total {
                        t.Error("Se esperaba 17 y se obtuvo", promediototal)
            }
            
            fmt.Println("Prueba de integracion del promedio total fue exitosa")
}
