/*
Archivo:      practica_01.go
Autor:        Edwin Yoner
Creado:       2025-02-16
Última mod:   2025-02-16
Licencia:     MIT
*/

// hola_mundo.go - Primer programa en Go
package main // Todo programa ejecutable debe pertenecer al paquete main

import "fmt" // Importar el paquete fmt para entrada/salida básica

// Ejemplo extendido
func main() {
	nombre := "Edwin Yoner"
	edad := 28
	fmt.Printf("¡Hola! Soy %s y tengo %d años\n", nombre, edad)
	fmt.Println(`Cadena multilínea:
    Línea 1
    Línea 2`)
	fmt.Println("Caracteres especiales: \t← Tabulación • \u2665 ← Unicode")
}
