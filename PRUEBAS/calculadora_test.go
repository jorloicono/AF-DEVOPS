package main

import "testing"
import "fmt"

func TestSuma(t *testing.T) {
	suma := Suma(7, 7)
	if suma != 14 {
		t.Error("Se esperaba un 14 y se ha obtenido ", suma)
	}

	fmt.Println("Prueba de suma exitosa")
}

func TestResta(t *testing.T) {
	resta := Resta(7, 7)
	if resta != 0 {
		t.Error("Se esperaba un 0 y se ha obtenido ", resta)
	}

	fmt.Println("Prueba de resta exitosa")
}

func TestMultiplicacion(t *testing.T) {
	Multiplicacion := Multiplicacion(7, 7)
	if Multiplicacion != 49 {
		t.Error("Se esperaba un 49 y se ha obtenido ", Multiplicacion)
	}

	fmt.Println("Prueba de Multiplicacion exitosa")
}

func TestDivision(t *testing.T) {
	Division := Division(7, 7)
	if Division != 1 {
		t.Error("Se esperaba un 1 y se ha obtenido ", Division)
	}

	fmt.Println("Prueba de Division exitosa")
}
