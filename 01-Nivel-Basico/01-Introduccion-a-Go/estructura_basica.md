# Estructura Básica de un Programa en Go

## 🏗 Componentes Fundamentales

### 1. Declaración de Paquete (`package`)
- **Primera línea obligatoria** en todo archivo `.go`
- Define a qué paquete pertenece el código
- `package main` para programas ejecutables

```go
package main  // Paquete principal para crear un ejecutable
```

### 2. Importación de Paquetes (`import`)
- Lista de paquetes externos necesarios
- Dos formatos válidos:

```go
// Formato agrupado
import (
    "fmt"
    "math/rand"
)

// Formato separado
import "os"
import "strings"
```

### 3. Función Principal (`func main()`)
- Punto de entrada del programa
- Obligatoria en paquetes `main`

```go
func main() {
    fmt.Println("¡Programa iniciado!")
}
```

## 📂 Organización del Código

### Estructura de un Módulo Básico
```
mi-proyecto/
├── go.mod       // Archivo de definición del módulo
├── main.go      // Punto de entrada principal
└── utils/
    └── math.go  // Paquete utilitario
```

### Visibilidad de Elementos
| Convención       | Visibilidad | Ejemplo              |
|------------------|-------------|----------------------|
| Nombre en mayúsc | Público     | `func Sumar() {...}` |
| Nombre en minúsc | Privado     | `func restar() {...}`|

```go
// En utils/math.go
package utils

// Función pública (exportada)
func Sumar(a, b int) int {
    return a + b
}

// Función privada (solo dentro del paquete)
func restar(a, b int) int {
    return a - b
}
```

## 🛠 Convenciones y Buenas Prácticas

### 1. **Orden de Elementos**
1. Declaración de paquete
2. Imports
3. Constantes/Variables globales
4. Funciones
5. Métodos
6. Structs

### 2. **Manejo de Errores**
- Convención de retornar `(resultado, error)`

```go
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("no se puede dividir por cero")
    }
    return a / b, nil
}
```

### 3. **Formateo Automático**
- Usar `gofmt` o `goimports` para:
    - Ordenar imports
    - Indentación consistente
    - Estilo canónico

```bash
   # Formatear archivo
   go fmt archivo.go
```

### 4. **Gestión de Dependencias**
- Usar módulos Go (`go.mod`)
- Ejemplo de `go.mod` básico:
```mod
module github.com/edwinyoner/specialization-go

go 1.21

require (
    github.com/gorilla/mux v1.8.0
)
```

## 🚨 Errores Comunes

### 1. **Paquete main sin función main**
```go
package main
// Falta func main() - Error: "runtime.main_main·f: function main is undeclared"
```

### 2. **Imports no utilizados**
```go
import "math"  // Error si no se usa
```

### 3. **Nombre de paquete incorrecto**
```go
// Archivo: calculadora.go
package main  // Error si es parte de un paquete utilitario
```

## � Ejemplo Completo

```go
package main

import (
    "fmt"
    "mi-proyecto/utils"
)

const Version = "1.0.0"

func main() {
    fmt.Println("Versión:", Version)
    
    suma := utils.Sumar(5, 3)
    fmt.Println("5 + 3 =", suma)
    
    _, err := dividir(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    }
}

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return a / b, nil
}
```
---

▶ **Siguiente**: [Hola Mundo](./hola_mundo.go)

---

▶ **Siguiente**: [Compilación y Ejecución](./go_run_build.md)

---