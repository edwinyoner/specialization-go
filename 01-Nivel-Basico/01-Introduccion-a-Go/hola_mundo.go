/*
Archivo:      hola_mundo.go
Autor:        Edwin Yoner
Creado:       2025-02-16
Última mod:   2025-02-16
Licencia:     MIT
*/

// hola_mundo.go - Primer programa en Go
package main // Todo programa ejecutable debe pertenecer al paquete main

import "fmt" // Importar el paquete fmt para entrada/salida básica

func main() {
	// La función main es el punto de entrada del programa
	// fmt.Println imprime texto en la consola y agrega un salto de línea

	fmt.Println("¡Hola, Mundo!") // Imprimir el clásico mensaje

	// Variaciones útiles para practicar:
	fmt.Print("Este mensaje no tiene salto de línea al final.")
	fmt.Printf("\nHola %s!\n", "Programador") // Formateo de cadenas

	// Desafíos adicionales:
	// 1. Imprime tu nombre y edad usando fmt.Printf
	// 2. Intenta imprimir caracteres especiales: \n (nueva línea), \t (tabulación)
	// 3. Crea un mensaje que ocupe múltiples líneas
}

/*
EJECUCIÓN Y COMPILACIÓN:

1. Ejecutar directamente:
   $ go run hola_mundo.go

2. Compilar y ejecutar:
   $ go build hola_mundo.go
   $ ./hola_mundo (Linux/Mac) o hola_mundo.exe (Windows)

CARACTERÍSTICAS DESTACADAS:
- Punto y coma automático: Go los inserta implícitamente
- Codificación UTF-8 nativa: Podemos usar caracteres especiales directamente
- Paquete main obligatorio para programas ejecutables
*/
